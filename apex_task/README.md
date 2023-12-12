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

I failed to get this part working. From what I could understand I needed to use the bridge (default) for the network.

> docker build -t apex:server -f Dockerfile . \
> docker build -t apex:client -f Dockerfile . \
> docker run -it --net=bridge apex:server \
> docker run -it --net=bridge apex:client

Before the above I simply tried to get the server working.
Exposing server port 8080 and using CMD for the client, I would get an EOF and then error connecting to 8080.