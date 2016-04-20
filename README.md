# go-react-boilerplate

A boilerplate for Golang / React Single Page Web Apps. Golang is the backend api and file-server. Reactjs is used for ONLY client-side views. Axios is used for xhr requests.

**Note: I use shell scripts. If you use windows, you'll have to write your own client/build.sh**


### Installation

```
# clone into your gopath
git clone https://github.com/ds0nt/go-react-boilerplate $GOPATH/src/my_project_name

cd $GOPATH/src/my_project_name

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

sudo npm install -g watchman browserify myth

# continually watch the clients css, js, and copy assets and build when changed
./build.sh

# build minified production version of the code.
# also, this will copy index.prod.html instead of index.html
NODE_ENV=production ./build.sh

```


### Notes on the code.

The server reads the config.toml file, then sets up a file-server and api-server accordingly. Just have a look at main.go. I left it all compacted into one file, so that you have the freedom to split it into files the way you want it to. You will also want to connect your own backend.

**Backend Recommendations**

 - [BoltDB](https://github.com/boltdb/bolt): is embedded, meaning that you don't actually have to install it. Great for tiny single-server apps. It's similar to redis.
 - [Mongo](http://labix.org/mgo): jf you want a document store.
 - [Redis](https://github.com/garyburd/redigo): if you want a key/value store.
 - [GORM](http://jinzhu.me/gorm/): if you want a relational database.


Mess around with the client/src/app.js, client/styles/app.css, client/public/index.html and client/public/index.prod.html

The client uses browserify to bundle up all of the application files into a single app.js. Similarily it uses myth to auto-prefix the css, and bundle it up into a single app.css file. This will happen automatically if you are running the ./build.sh file.


The client will use axios to make ajax requests. You might see that it uses a function apiPath(). This originates from the index.html file that is quickly rendered by the server with configuration before running.

The server api uses http://github.com/ant0ine/go-json-rest/rest, which is a thin layer around the standard "net/http" to provide nice json API functionality, and also has some nice middlewares.

Also, when running the server, you can do this:

```bash
./projectname -conf "production.toml"
```

Which is a nice way to have multiple configurations.

For production, I recommend using supervisorctl or systemd to keep the binary running, and to pipe the logs into somewhere like: /var/log/project.err.log and /var/log/project.log

Good luck!

### Contribute

 If you can pretty things up with your style of coding, please do so and pull request :)

 Cheers guys
