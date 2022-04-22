import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { ProposalsService } from 'src/app/services/proposals.service';

@Component({
  selector: 'app-proposals',
  templateUrl: './proposals.component.html',
  styleUrls: ['./proposals.component.css']
})
export class ProposalsComponent implements OnInit {



  constructor(public proposalSrvs: ProposalsService, public authSrvs: AuthService) {
    // this.proposeIdea()
    this.getPendingIdeas()
  }

  ngOnInit(): void {
  }

  proposeIdea(item: any) {
    this.proposalSrvs.proposeIdea(item)
  }

  getPendingIdeas() {
    this.proposalSrvs.loadPendingIdeas()
  }

  approve(item: any) {
    console.log(item);
    this.proposalSrvs.approveIdea(item)
  }

  disapprove(item: any) {
    console.log(item);
    this.proposalSrvs.disapproveIdea(item)
  }

}
