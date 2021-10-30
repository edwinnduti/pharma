import '../Styling/EntireLogins.css';

export const EntireLogins = () => {
    return(
        <ul>
            <li><a href="login.html" class="round green">Admin Login<span class="round">Caution Only for Admin!!!.</span></a></li>
            <li><a href="login.html" class="round red"> User Signin<span class="round">Take a look. This product is totally rad! </span></a></li>
            <li><a href="reg.html" class="round yellow">User Signup <span class="round">But only if you really, really want to!!!</span></a></li>
        </ul> 
    );
}