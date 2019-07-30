import React from 'react';
import {connect} from 'react-redux'
import {Layout, Header, Content, Grid, Cell, Button} from 'react-mdl';

const main = React.createClass({
  render: function () {
    return (
     <Layout fixedHeader>
        <Header title="React+Redux go">
        </Header>
        <Content className="mdl-color-text--grey-600">
          <Grid>
              {this.props.children}
         </Grid>
       </Content>
     </Layout>
    );
  }
});

//Map the local state directly to the state tree in the combined reducer.
export const AppLayout = connect((state) => state)(main);
