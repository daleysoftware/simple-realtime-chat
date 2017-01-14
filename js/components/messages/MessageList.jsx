let React = require('react');

let Message = require('./Message.jsx');

class MessageList extends React.Component{
    render(){
        return (
            <ul>{
                this.props.messages.map(message => {
                    return (
                        <Message key={message.id} message={message} />
                    )
                })
            }</ul>
        )
    }
}

MessageList.propTypes = {
    messages: React.PropTypes.array.isRequired
};

module.exports = MessageList;