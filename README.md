GoLang Terminal Messenger
=========================
GoLang Terminal Messenger is a client server messaging application developed using [gRPC](https://grpc.io/) framework in GoLang.

It involves the use of the gRPC Remote Procedure Call system, developed by Google and Protocol Buffers as its underlying message interchange format, also developed by Google. 

And because of the use of gRPC and Protocol Buffers, one can build the project using the proto file `proto/define.proto` which makes it a cross-platform messaging application.

Refer [gRPC](https://grpc.io/docs/) and [protobuf](https://developers.google.com/protocol-buffers/) for more understanding

Project Directories
===================
At a high level, there are 3 main directories, the **client**, **proto** and **server**

client
------
This contains the client implementation and it's setup in the file `client.go`

server
------
This contains the server implementation and along with it's configurations in the file `server.go`

proto
-----
Finally, the `proto` directory contains the structure, definitions and methods in the Google's Protocol Buffer format, which can be used to generate its equivalent in other languages, like I did in GoLang.

executables
-----------
This directory contains windows executable files of corresponding `client/client.go` and `server/server.go`

Installation and Setup
======================
1) Install and setup GoLang in your device https://golang.org/.
2) Now, run the following command:
        
        $ go get github.com/Shivam010/golang-terminal-messenger

Contributing
============
Changes and improvements are more than welcome! 
Feel free to fork and open a pull request. 
And Please make your changes in a specific branch and request to pull into master! If you can, please make sure the game fully works before sending the PR, as that will help speed up the process.

License
=======
Protocol Buffer to TypeScript Plugin is licensed under the [MIT license](https://github.com/Shivam010/golang-terminal-messenger/blob/master/LICENSE).
