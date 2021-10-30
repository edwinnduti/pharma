import '../Styling/Login.css';

export const Register = () => {
    return(
        <div class="login">
            <div class="login-triangle"></div>
  
            <h2 class="login-header">REGISTER</h2>

            <form class="login-container"  action="pregister.php" method="get">
                <p><input type="text" name="username" placeholder="username" required="required" /></p>
                
                <p><input type="password" name="password" placeholder="Password" required="required" /></p>
                
                <p><input type="password" name="cpwd" placeholder=" Confirm Password" required="required" /></p>
                
                <p><input type="text" name="mail" placeholder="EMAIL" required="required" /></p>
                
                <p><input type="text" name="phone" placeholder="phone"  maxlength="10" required="required" /></p>
                
                <p> <input type="submit" name="submit"   align="center" value="REGISTER" class="btn-login"/></p>
                
            </form>
        </div>
    );
}