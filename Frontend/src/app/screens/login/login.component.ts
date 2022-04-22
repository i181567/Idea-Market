import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private AuthSrvs: AuthService) { }

  ngOnInit(): void {
  }

  public login(usernameField:any, passwordField:any) {
    this.AuthSrvs.login(usernameField.value,passwordField.value)
  }

}
