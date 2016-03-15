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
    axios.get(apiPath('/apps'))
      .then(res => this.setState({ apps: res.data }))
      .catch(err => console.error(err))
  }

  componentWillUnmount() {

  }

  render() {

    return (
      <div>
        <h1>Forky Admin</h1>
        { JSON.stringify(this.state.apps) }
      </div>
    )
  }
}

ReactDOM.render((<App />), document.getElementById('application'))
