import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { IdeasService } from 'src/app/services/ideas.service';

@Component({
  selector: 'app-my-ideas',
  templateUrl: './my-ideas.component.html',
  styleUrls: ['./my-ideas.component.css']
})
export class MyIdeasComponent implements OnInit {
  constructor(public ideaService: IdeasService, public authSrvs: AuthService) { }

  ngOnInit(): void {
    this.ideaService.myIdeas()
  }

  startBidding(item: any) {
    this.ideaService.startBidding(item)
  }
  stopBidding(item: any) {
    this.ideaService.stopBidding(item)
  }
}
