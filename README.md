# go-react-boilerplate

A boilerplate for API driven single-page web applications with a react.js front-end

# Installation

Install golang. Setup your $GOPATH

```bash
# for example, add this to your .bashrc
export GOPATH="/home/MYUSER/go"

# then source it
. ~/.bashrc
```

```bash
# create the gopath if you havn't already :p
mkdir -p $GOPATH/src

# clone
git clone https://github.com/ds0nt/go-react-boilerplate $GOPATH/src/my_project_name

cd $GOPATH/src/my_project_name
```
Because go uses the src/ folder structure to include packages, find-replace *.go files "ds0nt.com" with "my_project_name".

Once you've done that...


### Server build steps

```bash
# install dependency packages
go get

# compile source into a binary
go build

# execute said binary (start the server)
./my_project_name
```


### Client build steps

```bash
cd client

# get the front-end dependencies
npm install

# continually watch the clients css, js, and copy assets and build when changed
./build.sh

```


### Notes on the code.

Mess around with the client/src/app.js, client/styles/app.css, and with config.go, config.toml, and server/server.go, server/routes/*

The client uses browserify to bundle up all of the application files into a single app.js. Similarily it uses myth to auto-prefix the css, and bundle it up into a single app.css file. This will happen automatically if you are running the ./build.sh file.

For the server, it first reads the config.toml file, then it will run client/render.go which will render out an index.html file with our configuration loaded in. This is a nice way to pass in the correct API Ports etc.

To add more routes, you can look at the /server/routes/apps.go. Its nice and simple and boring.

I do recommend using well structured responses, and I will add more examples soon!

```
GET http://localhost:4001/apps
```

The client will use axios to make ajax requests. You might see that it uses a function apiPath(). This originates from the index.html file that is quickly rendered by the server with configuration before running.
 
The server api uses http://github.com/ant0ine/go-json-rest/rest, which is a thin layer around the standard "net/http" to provide nice json API functionality, and also has some nice middlewares.

Also, when running the server, you can do this:

```bash
./projectname -conf "production.toml"
```

Which is a nice way to have multiple configurations.

For production, I recommend using supervisorctl to keep the binary running, and to pipe the logs into somewhere like: /var/log/project.err.log and /var/log/project.log

Good luck!


### Contribute
 
 If you can pretty things up with your style of coding, please do so and pull request :)
 
 Cheers guys

