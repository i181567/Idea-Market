import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HttpClientModule } from '@angular/common/http'

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavigatorComponent } from './shared/navigator/navigator.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FormsModule } from '@angular/forms';
import { HomeComponent } from './screens/home/home.component';
import { IdeasComponent } from './screens/ideas/ideas.component';
import { ProposalsComponent } from './screens/proposals/proposals.component';
import { AuctionComponent } from './screens/auction/auction.component';
import { LoginComponent } from './screens/login/login.component';
import { RegisterComponent } from './screens/register/register.component';
import { MyIdeasComponent } from './screens/my-ideas/my-ideas.component';
import { AccountComponent } from './screens/account/account.component';
import { ProposeIdeaComponent } from './screens/propose-idea/propose-idea.component';
import { FooterComponent } from './shared/footer/footer.component';
import { FourOhfourComponent } from './screens/four-ohfour/four-ohfour.component';

@NgModule({
  declarations: [
    AppComponent,
    NavigatorComponent,
    HomeComponent,
    IdeasComponent,
    ProposalsComponent,
    AuctionComponent,
    LoginComponent,
    RegisterComponent,
    MyIdeasComponent,
    AccountComponent,
    ProposeIdeaComponent,
    FooterComponent,
    FourOhfourComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    FormsModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
