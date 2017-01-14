let React = require('react');

let ChannelForm = require('./ChannelForm.jsx');
let ChannelList = require('./ChannelList.jsx');

class ChannelSection extends React.Component {
    render() {
        return (
            <div className="support panel panel-primary">
                <div className="panel-heading">
                     <strong>Channels</strong>
                </div>
                <div className="panel-body channels">
                    <ChannelList {...this.props} />
                    <ChannelForm {...this.props} />
                </div>
            </div>
        )
    }
}

ChannelSection.propTypes = {
    channels: React.PropTypes.array.isRequired,
    setChannel: React.PropTypes.func.isRequired,
    addChannel: React.PropTypes.func.isRequired,
    activeChannel: React.PropTypes.object
};

module.exports = ChannelSection;