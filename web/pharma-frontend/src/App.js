import {
  BrowserRouter as Router,
  
} from "react-router-dom";

import { Routes } from './Routes';


const App = () => {
  return (
    <div>
      <Router>
          <Routes />
      </Router>


    </div>
  );
}

export default App;
