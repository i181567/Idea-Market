import { IdeasComponent } from './screens/ideas/ideas.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './screens/home/home.component';
import { ProposalsComponent } from './screens/proposals/proposals.component';
import { AuctionComponent } from './screens/auction/auction.component';
import { RegisterComponent } from './screens/register/register.component';
import { LoginComponent } from './screens/login/login.component';
import { MyIdeasComponent } from './screens/my-ideas/my-ideas.component';
import { AccountComponent } from './screens/account/account.component';
import { ProposeIdeaComponent } from './screens/propose-idea/propose-idea.component';
import { FourOhfourComponent } from './screens/four-ohfour/four-ohfour.component';

const routes: Routes = [
  { path: "", redirectTo: "login", pathMatch:"full" },
  { path: "home", component: HomeComponent },
  { path: "signup", component: RegisterComponent },
  { path: "login", component: LoginComponent },
  { path: "ideas", component: IdeasComponent },
  { path: "proposals", component: ProposalsComponent },
  { path: "auction", component: AuctionComponent },
  { path: "account", component: AccountComponent },
  { path: "my-ideas", component: MyIdeasComponent },
  { path: "propose-new", component: ProposeIdeaComponent },
  { path: "**", component: FourOhfourComponent },

];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
