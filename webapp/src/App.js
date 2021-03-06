import React, {Component} from 'react';
import Header from "./components/Header";
import Home from "./components/Home";
import {formatError, parseRoute} from "./components/Utils";
import BuildLog from "./components/BuildLog";
import Footer from "./components/Footer";
import api from "./components/api";
import Settings from "./components/Settings";

class App extends Component {
    constructor(props) {
        super(props)
        this.state = {
            system: {cpu:0, memory:0, disk:0},
            environment: {version:'0.1.0', arch:'arm64'},
        }
    }

    componentDidMount() {
        this.getSystemEnvironment()
        this.getSystemMonitor()
    }

    poll = () => {
        // Polls every 2s
        setTimeout(this.getSystemMonitor.bind(this), 2000);
    }

    getSystemEnvironment() {
        api.systemEnvironment().then(response => {
            this.setState({environment: response.data.record})
        })
        .catch(e => {
            console.log(formatError(e.response.data))
            this.setState({error: formatError(e.response.data), message: ''});
        })
    }

    getSystemMonitor() {
        api.systemMonitor().then(response => {
            this.setState({system: response.data.record})
        })
        .catch(e => {
            console.log(formatError(e.response.data))
            this.setState({error: formatError(e.response.data), message: ''});
        })
        .finally( ()=> {
            this.poll()
        })
    }

    render() {
        const r = parseRoute()

        return (
            <div>
                <Header environment={this.state.environment}/>

                {r.section===''? <Home/> : ''}
                {r.section==='builds'? <BuildLog buildId={r.sectionId} /> : ''}
                {r.section==='settings'? <Settings /> : ''}

                <Footer system={this.state.system} />
            </div>
        );
    }
}

export default App;
