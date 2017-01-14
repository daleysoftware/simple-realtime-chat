let React = require('react');
let ReactDOM = require('react-dom');

window.React = React;

let App = require('./components/App.jsx');

ReactDOM.render(<App />, document.getElementById('app'));