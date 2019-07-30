import React from 'react';
import {Textfield, Button, Spinner} from 'react-mdl';
import {setDataFieldValue, fetchData} from '../state/api-data'

export const FetchDataForm = React.createClass({
  handleClick: function() {
    this.props.dispatch(fetchData(this.props.Data.get("ip")))
  },
  setField: function(e) {
    this.props.dispatch(setDataFieldValue(e.target.name, e.target.value))
  },
  render: function() {
    return (
      <div>
        <Textfield
          label="Enter IP"
          floatingLabel
          name="ip"
          onChange={this.setField}
          value={this.props.Data.get("ip")}/>
        <Button onClick={this.handleClick} colored raised>
          Ping
          {this.props.Data.toJS().Loading === true ? <Spinner /> : null }
        </Button>
      </div>
    )
  }
})
