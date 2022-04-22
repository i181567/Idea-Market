import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { apiip } from '../serverConfig';


@Injectable({
  providedIn: 'root'
})
export class AuthService {

  public nullVal = {
    Balance: 0,
    Email: "u7",
    Password: "p7",
    Phonenumber: "u7",
    Username: "u7"
  }

  public activeUser: any = {
    Balance: 0,
    Email: "",
    Password: "",
    Phonenumber: "",
    Username: ""
  };

  constructor(private http: HttpClient, private router: Router) { }

  public logout() {
    this.router.navigateByUrl('/login')
    this.activeUser = {
      Balance: 0,
      Email: "",
      Password: "",
      Phonenumber: "",
      Username: ""
    };
  }

  public register(
    username: string, email: string, password: string, telephone: string
  ) {
    // Username, Password, Email, Phonenumber
    this.http.post(`${apiip}/signup`, {
      Username: username,
      Password: password,
      Email: email,
      Phonenumber: telephone
    })
      .toPromise()
      .then(res => {
        if (res === "OK") {
          this.router.navigateByUrl("/login")
        } else {
          alert("Choose a different username and email")
        }
      })
      .catch(err => {
        console.log(err);
      })
  }

  public login(username: string, password: string) {
    this.http.post<any>(`${apiip}/signin`, {
      Username: username,
      Password: password
    })
      .toPromise()
      .then(res => {
        if (res !== "ERR") {
          this.activeUser = res[0]
          console.log(this.activeUser);
          this.router.navigateByUrl("/ideas")
        } else {
          alert(res)
        }
      })
      .catch(err => {
        console.log(err);
      })
  }

  public loginNoRedirect(username: string, password: string) {
    this.http.post<any>(`${apiip}/signin`, {
      Username: username,
      Password: password
    })
      .toPromise()
      .then(res => {
        if (res !== "ERR") {
          this.activeUser = res[0]
          console.log(this.activeUser);
        } else {
          alert(res)
        }
      })
      .catch(err => {
        console.log(err);
      })
  }

  addCoins(coins: number, price: number) {
    console.log(this.activeUser.Username);
    console.log(this.activeUser.Password);
    console.log(price);
    console.log(coins);

    this.http.post(`${apiip}/deposit`, {
      Username: this.activeUser.Username,
      Password: this.activeUser.Password,
      Balance: coins
    })
      .toPromise()
      .then(res => {
        console.log(res);

      })
      .catch(err => {
        console.log(err);
      })
  }

  updateUser(updatedUser: any) {
    console.log(updatedUser);
    this.http.post(`${apiip}/updateuser`,
      {
        Username: updatedUser.Username,
        Password: updatedUser.Password,
        Email: updatedUser.Email,
        Phonenumber: updatedUser.Phonenumber
      })
      .toPromise()
      .then(res => {
        console.log(res);
        alert("You'll be logged out now")
        this.router.navigateByUrl("/login")
      })
      .catch(err => {
        console.log(err)
      })
  }

  deleteUser(updatedUser: any) {
    console.log(updatedUser);
    this.http.post(`${apiip}/deleteuser`,
      {
        Username: updatedUser.Username,
        Password: updatedUser.Password
      })
      .toPromise()
      .then(res => {
        console.log(res);
        alert("You'll be logged out now")
        this.router.navigateByUrl("/login")
      })
      .catch(err => {
        console.log(err)
      })
  }



}