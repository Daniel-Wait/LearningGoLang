# Task Report

## Chatroom

I used TCP as the communications protocol.
The server saves client connections in a Go Map and each client is forwarded
messages which are not its own. \
The client expects console text input and reads broadcast messages from other addresses.

## Local Testing

I tested my code in Windows Command Prompt (CMD) as I extended the program functionality. \
Call "go run ." in server and client subfolders respectively to start a session and test \
The image ___apex_task\windows_3xClients_test.png___ shows the server interacting with 3 clients each sending test messages.

## Docker Container Communication

Start the server first on the Bridge Network such that it shall default to IP address 172.17.0.2

In terminal #1
> cd server; docker build -t apex:server -f Dockerfile . \
> docker run -it --net=bridge apex:server \

In terminal #2
> cd client; docker build -t apex:client -f Dockerfile . \
> docker run -it --net=bridge apex:client
