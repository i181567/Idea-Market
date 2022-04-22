import { Component, OnInit } from '@angular/core';
import { AuctionService } from 'src/app/services/auction.service';
import { AuthService } from 'src/app/services/auth.service';
import { IdeasService } from 'src/app/services/ideas.service';

@Component({
  selector: 'app-ideas',
  templateUrl: './ideas.component.html',
  styleUrls: ['./ideas.component.css']
})
export class IdeasComponent implements OnInit {

  constructor(public ideaService: IdeasService, public authSrvs: AuthService, public auctionSrvs: AuctionService) {
  }

  ngOnInit(): void {
    this.ideaService.loadIdeas();
  }

  payToSee(ideaTitle: string) {
    this.ideaService.payToSee(this.authSrvs.activeUser.Username, ideaTitle);
    this.ideaService.loadIdeas()
  }

  isViewable(idea: any) {
    return (idea.Can_be_viewed_by.filter((vu: string) => vu !== "").findIndex((vu: string) => vu === this.authSrvs.activeUser.Username) !== -1)
    // this.authSrvs.activeUser.Username
  }


}
