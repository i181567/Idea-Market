import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  public password: string = "";
  public vPassword: string = "";

  constructor(public authSrvs: AuthService) {
    // this.authSrvs.register();
    // this.authSrvs.login();
  }

  ngOnInit(): void {
  }

  register(
    usernameField: any,
    emailField: any,
    passwordField: any,
    telephoneField: any
  ) {
    if (this.password !== this.vPassword) {
      return
    }

    this.authSrvs.register(usernameField.value, emailField.value, passwordField.value, telephoneField.value);

  }
}
