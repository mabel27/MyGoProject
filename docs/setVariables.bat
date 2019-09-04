@echo off

rem This batch create directory and set the variables need it.

ECHO Begin Set in variables...

if not exist "C:\Project\myExchangeProject" mkdir C:\Project\myExchangeProject

SET GOPATH=C:\Project\myExchangeProject
SET SLACKTOKEN=xoxp-735488424065-735488424737-744894194389-2fd39509c7061e8face78c8e14cdfd87
SET EXCHANGETOKEN=b2e3d360a5c775a403d9ddff35e33cbd

ECHO Done...

EXIT /B 
