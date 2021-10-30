import {
    Route,
    Switch,
    Redirect,
} from 'react-router-dom';
import { Login } from './pages/Login';
import { Register } from './pages/Register';
import { EntireLogins } from './pages/EntireLogins';

export const Routes = () => {
    return(
        <Switch>
            <Route basename={process.env.PUBLIC_URL} exact path="/" component={Login} />
            <Redirect exact from="/clinic" to="/" />
            <Route basename={process.env.PUBLIC_URL} exact path="/signup" component={Register} />
            <Route basename={process.env.PUBLIC_URL} exact path="/logins" component={EntireLogins} />
        </Switch>
    );
}