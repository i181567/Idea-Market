import { Component, OnInit } from '@angular/core';
import { AuctionService } from 'src/app/services/auction.service';
import { AuthService } from 'src/app/services/auth.service';
import { IdeasService } from 'src/app/services/ideas.service';

@Component({
  selector: 'app-auction',
  templateUrl: './auction.component.html',
  styleUrls: ['./auction.component.css']
})
export class AuctionComponent implements OnInit {

  public biddingPrice: number = 0;
  constructor(public authSrvs: AuthService, public auctionService: AuctionService, public ideaService: IdeasService) { }

  ngOnInit(): void {
    this.auctionService.loadAuctionables()
  }

  bid(ideatitle: string, biddingPrice: number) {
    this.auctionService.bid(ideatitle, this.authSrvs.activeUser.Username, biddingPrice);
  }

}
