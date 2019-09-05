@echo off

rem This batch create directory and set the variables need it.

ECHO Begin Set in variables...

if not exist "C:\Project\myExchangeProject" mkdir C:\Project\myExchangeProject

SET GOPATH=C:\Project\myExchangeProject
SET SLACKTOKEN=TOKEN HERE
SET EXCHANGETOKEN=TOKEN HERE

ECHO Done...

EXIT /B 
