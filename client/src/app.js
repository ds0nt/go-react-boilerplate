import React, { Component, PropTypes } from 'react'
import ReactDOM from 'react-dom'
import axios from 'axios'

class App extends Component {
  constructor() {
    this.state = {
      apps: []
    }
  }

  componentDidMount() {
    axios.get(API_URL + '/apps')
      .then(res => this.setState({ apps: res.data }))
      .catch(err => console.error(err))
  }

  componentWillUnmount() {

  }

  render() {

    return (
      <div>
        <h1>Go React Boilerplate</h1>
        <p>
          To build the production version of the client try: <br/>
          <b>NODE_ENV=production ./build.sh </b> <br/><br/>
          see build.sh for details
        </p>
        <h3>Apps</h3>
        <a href={API_URL + "/apps"}>GET /apps</a>
        <pre>
          { JSON.stringify(this.state.apps, false, 4) }
        </pre>
      </div>
    )
  }
}

ReactDOM.render((<App />), document.getElementById('application'))
