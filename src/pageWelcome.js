import React from 'react'

import {connect} from 'react-redux'

import {Grid, Cell} from 'react-mdl';
import {Link} from 'react-router'

const main = React.createClass({
  render: function() {
    return (
        <Cell col={12}>
          <h1>Welcome!</h1>
          <h2>Ping</h2>
          <Link to="ping" href="/ping">ping to host</Link>
        </Cell>
    )
  }
});

//Map the local state directly to the state tree in the combined reducer.
export const WelcomePage = connect((state) => state)(main);
