import React from "react";
import axios from "axios";

export default class SimpleForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      message: ""
    }
    this.getApi = this.getApi.bind(this)
    this.postApi = this.postApi.bind(this)
  }
  render() {
    return (
      <>
        <button onClick={this.getApi}>GET API</button>
        <button onClick={this.postApi}>POST API</button>
      </>
    )
  }
  getApi() {
    axios.get("http://127.0.0.1:3333/test")
    .then((response) => console.log(response))
    .catch(error => console.log(error))
  }
  postApi() {
    axios.post("http://127.0.0.1:3333/test")
    .then((response) => console.log(response))
    .catch(error => console.log(error))
  }
}