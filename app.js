let React = require('react');
let ReactDOM = require('react-dom');

window.React = React;

class Channel extends React.Component {

    onClick() {
        console.log(this.props.name);
    }

    render() {
        return (
            <li onClick={this.onClick.bind(this)}>{this.props.name}</li>
        )
    }
}

class ChannelList extends React.Component {
    render() {
        let channelComponents = this.props.channels.map(function(channel)
        {
            return (
                <Channel key={channel.name} name={channel.name} />
            );
        });

        return (
            <ul>
                {channelComponents}
            </ul>
        )
    }
}

class ChannelForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            channelName: ''
        };
    }

    onSubmit(e) {
        let channelName = this.state.channelName;
        this.setState({
            channelName: ''
        });
        this.props.addChannel(channelName);
        e.preventDefault();
    }

    onChange(e) {
        this.setState({
            channelName: e.target.value
        });
    }

    render() {
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <input
                    type="text"
                    onChange={this.onChange.bind(this)}
                    value={this.state.channelName} />
            </form>
        )
    }
}

class ChannelSection extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            channels: [
                {name: 'Hardware support'},
                {name: 'Software support'}
            ]
        };
    }

    addChannel(name) {
        let channels = this.state.channels;
        channels.push({name});
        this.setState({
            channels: channels
        });
    }

    render() {
        return (
            <div>
                <ChannelList channels={this.state.channels} />
                <ChannelForm addChannel={this.addChannel.bind(this)} />
            </div>
        )
    }
}

ReactDOM.render(<ChannelSection />, document.getElementById('app'));