let React = require('react');

class UserForm extends React.Component{
    onSubmit(e){
        e.preventDefault();
        const node = this.refs.userName;
        const userName = node.value;
        this.props.setUserName(userName);
        node.value = '';
    }

    render(){
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <div className='form-group'>
                    <input
                        ref='userName'
                        type='text'
                        className='form-control'
                        placeholder='Set Your Name...' />
                </div>
            </form>
        )
    }
}

UserForm.propTypes = {
    setUserName: React.PropTypes.func.isRequired
};

module.exports = UserForm;