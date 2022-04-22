import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent implements OnInit {

  public exchangeRate = 10;
  public coinsToAdd: number = 0;
  public emailAddress = ""

  public newEmail: string = this.authSrvs.activeUser.Email;
  public newPassword: string = this.authSrvs.activeUser.Password;
  public newPhonenumber: string = this.authSrvs.activeUser.Phonenumber;

  constructor(public authSrvs: AuthService, private router: Router) { }

  ngOnInit(): void {
    let asterisks = "*************************"
    this.emailAddress = asterisks.substring(0, this.authSrvs.activeUser.Email.length / 2) + this.authSrvs.activeUser.Email.substring(this.authSrvs.activeUser.Email.length / 2, this.authSrvs.activeUser.Email.length)
  }

  getPrice() {
    return this.exchangeRate * this.coinsToAdd;
  }

  addCoins() {
    this.authSrvs.addCoins(this.coinsToAdd, this.getPrice())
    this.authSrvs.loginNoRedirect(this.authSrvs.activeUser.Username, this.authSrvs.activeUser.Password);
  }

  updateUser() {
    this.authSrvs.updateUser({
      Username: this.authSrvs.activeUser.Username,
      Email: this.newEmail,
      Password: this.newPassword,
      Phonenumber: this.newPassword
    });
  }

  deleteUser() {
    this.authSrvs.deleteUser({
      Username: this.authSrvs.activeUser.Username,
      Email: this.newEmail,
      Password: this.newPassword,
      Phonenumber: this.newPassword
    });
  }


}
