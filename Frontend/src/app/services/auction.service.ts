import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { apiip } from '../serverConfig';
import { AuthService } from './auth.service';
import { IdeasService } from './ideas.service';

@Injectable({
  providedIn: 'root'
})
export class AuctionService {

  public ideas: any = []

  constructor(private http: HttpClient, private ideasService: IdeasService, private authService: AuthService) {
    this.loadAuctionables();
  }

  loadAuctionables() {
    this.http.post(`${apiip}/ideasinauction`, {
      Username: this.authService.activeUser.Username
    })
      .toPromise()
      .then(res => {
        console.log("Auctionables");
        console.log(res);
        this.ideas = res
      })
      .catch(res => {
        console.log(res);
      })
  }

  // Title        string
  // Username     string
  // Biddingprice float64
  bid(ideatitle: string, username: string, biddingPrice: number) {
    this.http.post(`${apiip}/bid`, {
      Title: ideatitle,
      Username: username,
      Biddingprice: biddingPrice
    })
      .toPromise()
      .then(res => {
        console.log(res);
        // this.ideasService.loadIdeas();
        this.ideasService.loadIdeas()
        this.ideasService.myIdeas()
        this.loadAuctionables();
      })
      .catch(err => {
        console.log(err);
      })
  }
}
