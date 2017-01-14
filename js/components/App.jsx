let React = require('react');

let ChannelSection = require('./channels/ChannelSection.jsx');

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            channels: []
        };
    }

    addChannel(name) {
        let {channels} = this.state;
        channels.push({id: channels.length, name});
        this.setState({channels});
        // TODO send to server
    }

    setChannel(activeChannel) {
        this.setState({activeChannel});
        // TODO get latest channel messages from server
    }

    render() {
        return (
            <div className="app">
                <div className="nav">
                    <ChannelSection
                        {...this.state}
                        addChannel={this.addChannel.bind(this)}
                        setChannel={this.setChannel.bind(this)}
                    />
                </div>
            </div>
        );
    }
}

module.exports = App;