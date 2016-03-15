#!/bin/bash

trap killgroup SIGINT

killgroup(){
  echo "killing..."
  kill 0
}

[ -d 'dist' ] || mkdir -p 'dist'

watchman public 'cp -urv public/* dist/' &
watchman src 'browserify -d -e src/app.js -t babelify -o dist/app.js' &
watchman styles 'myth styles/app.css dist/app.css' &

wait
