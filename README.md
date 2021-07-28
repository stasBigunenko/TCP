TCP-Server
Simple TCP server and client.

Usage:
This will create a small server and client side. Client can write a message and the server will respond.

At the moment the server knows the following commands: Hi, Bye, Time, Id, Help. Other commands are under developing.

The server automatically will connect with tcp connection on localhost (tcp on 127.0.0.1:8080). If you want to change any of these values you can use enviroment variables "PROTOCOL", "PORT" and/or "HOST".

The server will save all the commands from client side and also will save the unknown commands from client side in separate file. It was done with hope to improve the server's response with the most frequently asked commands which unknown.

