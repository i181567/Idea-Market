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

  public totalNumUser:number = 0

  constructor(private http: HttpClient, private router: Router) {

    if (this.checkPersist()) {
      this.loginPersistent()
    }
    this.totalUser()
  }

  public logout() {
    this.rmpersist()
    this.router.navigateByUrl('/login')
    this.activeUser = {
      Balance: 0,
      Email: "",
      Password: "",
      Phonenumber: "",
      Username: ""
    };
  }

  persist(username: string, password: string) {
    localStorage.setItem("U", username)
    localStorage.setItem("P", password)
  }
  rmpersist() {
    localStorage.removeItem("U")
    localStorage.removeItem("P")
  }
  checkPersist() {
    if (localStorage.getItem("U") !== null && localStorage.getItem("P") !== null) {
      return true
    }
    return false
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
          this.totalUser()
        } else {
          alert("Choose a different username and email")
        }
      })
      .catch(err => {
        console.log("Can't sign you up");
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
          this.persist(username, password)
          this.router.navigateByUrl("/ideas")

        } else {
          alert("Invalid User")
        }
      })
      .catch(err => {
        alert("Invalid User")
      })
  }

  public loginPersistent() {
    this.http.post<any>(`${apiip}/signin`, {
      Username: localStorage.getItem("U"),
      Password: localStorage.getItem("P")
    })
      .toPromise()
      .then(res => {
        if (res !== "ERR") {
          this.activeUser = res[0]
          console.log(this.activeUser);
          this.router.navigateByUrl("/ideas")

        } else {
          alert("Invalid User")
        }
      })
      .catch(err => {
        console.log(err);
        alert("Invalid User")
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
          alert("Invalid User")
        }
      })
      .catch(err => {
        alert("Invalid User")
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
        this.totalUser()
        this.router.navigateByUrl("/login")
      })
      .catch(err => {
        console.log(err)
      })
  }

  totalUser() {
    this.http.get<number>(`${apiip}/totalusers`)
      .toPromise()
      .then(res => {
        console.log(res);
        if (res !== undefined)
          this.totalNumUser = res
      })
      .catch(err => {
        console.log(err);
      })
  }



}