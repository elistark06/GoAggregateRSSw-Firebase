@echo off
:loop
go run main.go
echo Server stopped, restarting in 5 seconds...
timeout /t 0
goto loop