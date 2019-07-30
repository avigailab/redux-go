import React from 'react'

import {connect} from 'react-redux'

import {Grid, Cell} from 'react-mdl';

import {FetchDataForm} from './components/fetch-data'

const main = React.createClass({
  render: function() {
    /*let data = this.props.Data.get("Data").map(x => {
      return (
        <li>{x}</li>
      )
    })*/
    let data = this.props.Data.get("Data");
    return (
      <Grid>
        <Cell col={12}>
          <h1>Test Page</h1>
          <FetchDataForm {...this.props} />
          <h5>Has ping:</h5>
          <ol>
            {data === true ? 'True' : 'False'}
          </ol>
        </Cell>
      </Grid>
    )
  }
});

//Map the local state directly to the state tree in the combined reducer.
export const PingPage = connect((state) => state)(main);
