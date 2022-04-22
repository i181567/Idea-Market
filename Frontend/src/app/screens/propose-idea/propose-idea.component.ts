import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { ProposalsService } from 'src/app/services/proposals.service';

@Component({
  selector: 'app-propose-idea',
  templateUrl: './propose-idea.component.html',
  styleUrls: ['./propose-idea.component.css']
})
export class ProposeIdeaComponent implements OnInit {
  public nullValue = {
    Bidding: false,
    Can_be_viewed_by: [],
    Description: "",
    Domain: "",
    End_bidding_time: "",
    Hash_of_idea: "",
    Highest_bidder: "",
    Highest_bidding_price: "",
    Owners: [""],
    Ownership_price: 0,
    Pricing_history: [0],
    Problem: "",
    Start_bidding_time: "",
    Technologies_used: [""],
    Title: "",
    Viewing_price: 0
  }

  public newIdea = JSON.parse(JSON.stringify(this.nullValue));
  public newStack: string = ""
  constructor(public prpSrvs: ProposalsService, private authErvs: AuthService) { }

  ngOnInit(): void {
    this.checkStackDirt()
  }

  checkStackDirt() {
    let index = this.newIdea.Technologies_used.findIndex((ts: string) => ts === "");
    if (index !== -1) {
      this.newIdea.Technologies_used.splice(index, 1)
    }
  }

  addTechStack() {
    let index = this.newIdea.Technologies_used.findIndex((ts: string) => ts === "");
    if (index !== -1) {
      this.newIdea.Technologies_used.splice(index, 1)
    }

    let index2 = this.newIdea.Technologies_used.findIndex((ts: string) => ts === this.newStack);
    if (index2 !== -1) {
      return
    }
    if (this.newStack !== undefined && this.newStack !== null && this.newStack) {
      this.newIdea.Technologies_used.push(this.newStack);
      this.newStack = ""
    }
    console.log(this.newIdea.Technologies_used);
  }

  RemoveTechStack(i: number) {
    this.newIdea.Technologies_used.splice(i, 1);
  }

  proposeIdea() {
    this.newIdea.Owners = [this.authErvs.activeUser.Username]
    this.prpSrvs.proposeIdea(JSON.parse(JSON.stringify(this.newIdea)));
    this.newIdea = {
      Bidding: false,
      Can_be_viewed_by: [],
      Description: "",
      Domain: "",
      End_bidding_time: "",
      Hash_of_idea: "",
      Highest_bidder: "",
      Highest_bidding_price: "",
      Owners: [""],
      Ownership_price: 0,
      Pricing_history: [0],
      Problem: "",
      Start_bidding_time: "",
      Technologies_used: [""],
      Title: "",
      Viewing_price: 0
    }
    this.newStack = ""
    this.checkStackDirt()

  }

}
