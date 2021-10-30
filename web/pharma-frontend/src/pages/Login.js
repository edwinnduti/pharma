import '../Styling/Login.css';

export const Login = () => {
    return(
        <div class="login">
            <div class="login-triangle"></div>
            
            <h2 class="login-header">Log in</h2>

            <form class="login-container" method="post" action="log2.php">
                <p><input type="text" name="username" placeholder="username" required="required" /></p>
                <p><input type="password" name="password" placeholder="Password" required="required" /></p>
                <p> <input type="submit" name="submit"   align="center" value="LOGIN" class="btn-login"/></p>
            </form>
        </div>
    );
}