import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-navigator',
  templateUrl: './navigator.component.html',
  styleUrls: ['./navigator.component.css']
})
export class NavigatorComponent implements OnInit {
  public isMenuCollapsed = true;
  constructor(public auth: AuthService,) { }

  ngOnInit(): void {
  }
  logout() {
    this.auth.logout()
  }

  showHome() {
    let home = document.getElementById("Home")
    let root = document.getElementById("Root")
    if (home != undefined && root != undefined) {
      root.setAttribute('style', 'height:0pt;overflow:hidden')
      home.setAttribute('style', 'height:max-content;overflow:hidden')
    }
  }
}
