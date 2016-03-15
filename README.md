# go-react-boilerplate

A boilerplate for API driven web applications with a react.js front-end

# Installation

 - Clone the repository into a folder like ```/mygopath/src/projectname```
 - Because the boilerplate was made for ```/mygopath/src/ds0nt.com``` you need to ctrl+f and replace all 'ds0nt.com's with 'projectname'

# Usage

Building the client. Have npm and nodejs installed.

```bash
cd client

# get the front-end dependencies
npm install

# continually watch the clients css, js, and copy assets and build when changed
./build.sh

```

Running the server

```bash
go build .
./myproject
```

# Development

It's pretty straightforward. Mess around with the client/src/app.js, client/styles/app.css, and with config.go, config.toml, and server/server.go, server/routes/*

There is one example route
```
GET http://localhost:4001/apps
```
 will hit up the golang api provided
 
 # Contribute
 
 If you can pretty things up with your style of coding, please do so and pull request :)
 
 Cheers guys

