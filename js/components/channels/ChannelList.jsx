let React = require('react');

let Channel = require('./Channel.jsx');

class ChannelList extends React.Component {
    render() {
        return (
            <ul>{
                this.props.channels.map(channel => {
                    return <Channel
                        key={channel.id}
                        channel={channel}
                        {...this.props}
                    />
                })
            }</ul>
        );
    }
}

ChannelList.propTypes = {
    channels: React.PropTypes.array.isRequired,
    setChannel: React.PropTypes.func.isRequired,
    activeChannel: React.PropTypes.object
};

module.exports = ChannelList;