import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { apiip } from '../serverConfig';
import { AuctionService } from './auction.service';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class IdeasService {

  public ideas: any[] = []
  public myideas: any[] = []

  constructor(private http: HttpClient, private authSrvs: AuthService) {
    this.loadIdeas()
    this.myIdeas()
  }

  public loadIdeas() {
    this.http.post<any>(`${apiip}/showideas`,
      {
        Username: this.authSrvs.activeUser.Username
      }, { responseType: "json" })
      .toPromise()
      .then(res => {
        // console.log(res);

        if (res !== undefined) {
          this.ideas = res
        }
        console.log("showideas");

        console.log(this.ideas);
      })
      .catch(err => {
        console.log(err);
      })
  }

  public myIdeas() {
    this.http.post<any>(`${apiip}/myideas`,
      {
        "Username": this.authSrvs.activeUser.Username
      }, { responseType: 'json' })
      .toPromise()
      .then(res => {
        if (res !== undefined) {
          this.myideas = res
          console.log("myideas");
          console.log(this.myideas);
        }
      })
      .catch(err => {
        console.log(err);
      })
  }

  public startBidding(idea: any) {
    console.log(idea);
    this.http.post(`${apiip}/startbidding`, {
      Username: idea.Username,
      Title: idea.Title,
      Password: idea.Password,
    })
      .toPromise()
      .then(res => {
        this.myIdeas();
        this.loadIdeas();
      })
      .catch(err => {
        alert("Could not start bidding")
      })
  }

  public stopBidding(idea: any) {
    console.log(idea);
    this.http.post(`${apiip}/stopbidding`, idea)
      .toPromise()
      .then(res => {
        this.myIdeas();
        this.loadIdeas();
      })
      .catch(err => {
        alert("Could not stop bidding")
      })
  }

  public payToSee(username: string, ideaTitle: string) {
    console.log(username);
    console.log(ideaTitle);

    this.http.post(`${apiip}/viewidea`, {
      Username: username,
      Title: ideaTitle
    })
      .toPromise()
      .then(res => {
        this.myIdeas();
        this.loadIdeas();
      })
      .catch(err => {
        console.log(err);
      })
  }

}
