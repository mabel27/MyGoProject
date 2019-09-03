package main

import (
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  //"log"
  "net/http"

  "github.com/nlopes/slack"
  "github.com/tidwall/gjson"

)

func getenv(name string) string {
  v := os.Getenv(name)
  if v == "" {
    panic("missing required environment variable " + name)
  }
  return v
}

func main() {
  token := getenv("SLACKTOKEN")
  api := slack.New(token)
  rtm := api.NewRTM()
  go rtm.ManageConnection()

  Loop:
  for {
    select {
    case msg := <-rtm.IncomingEvents:
      fmt.Print("Event Received: ")
      switch ev := msg.Data.(type) {

        case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					botRespond(rtm, ev, prefix)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
      }
    }
  }

  func botRespond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)

  response =  getExchange(text)
  rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
}

//Call get exchange rate API
func getExchange(currency string) string {
  var positive string
  var negative string

  response, err := http.Get("http://www.apilayer.net/api/live?access_key=b2e3d360a5c775a403d9ddff35e33cbd&format=1")
  if err != nil {
      fmt.Print(err.Error())
      os.Exit(1)
  }
  responseData, err := ioutil.ReadAll(response.Body)
  value := gjson.GetBytes(responseData, "quotes."+currency+"")

  positive = ":dollar: The exchange rate from " + currency + " = " +value.String()
  negative = "Please check your currency format:exclamation:"

  if gjson.GetBytes(responseData,  "quotes."+currency+"").Exists() {
  	return (positive)
  }else
  {
    return(negative)
  }
}
