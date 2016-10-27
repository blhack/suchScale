var $ = require("jquery");
import React from "react";
import ReactDOM from "react-dom";


class Story extends React.Component {
	constructor(props) {
		super(props);
		}

	render() {
		return(
			<h1><a href={this.props.Url}>{this.props.Title}</a></h1>
		);
	}
}

class App extends React.Component {

	constructor(props) {
		super(props);
		this.state = {stories:[]};
	}

	componentDidMount() {
		console.log("Component did mount");
		window.getData = this.getData;
		this.getData();
		this.timerId = setInterval(() => this.getData(), 5000);
	}

	getData() {
		console.log("Getting data");
		$.get("http://localhost:3000/stories/", (data) => {
			console.log(data);
			this.setState({stories:data});
		});
	}

	render() {
		return(
			<div>
				{this.state.stories.map(function(story, index) {
					return(
							<Story {...story} key={index} />
						)
				})}
			</div>
		);
	}
}


function main() {
	ReactDOM.render(
		<App />,
		document.getElementById('root')
	);
}

window.main = main;