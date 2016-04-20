#!/bin/bash

# ctrl c kills all the processes
trap killgroup SIGINT
killgroup(){
  echo "killing..."
  kill 0
}

[ -d 'dist' ] || mkdir -p 'dist'

if [[ "${NODE_ENV}" == "production" ]]; then
  cp -urv public/* dist/
  browserify -g uglifyify -e src/app.js -t babelify | uglifyjs -c > dist/app.js
  myth styles/app.css dist/app.css
  rm dist/index.html
  mv dist/index.prod.html dist/index.html
else
  watchman public 'cp -urv public/* dist/' &
  watchman src 'browserify -d -e src/app.js -t babelify -o dist/app.js' &
  watchman styles 'myth styles/app.css dist/app.css' &
fi

wait
