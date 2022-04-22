import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { apiip, nlpapi } from '../serverConfig';

@Injectable({
  providedIn: 'root'
})
export class ProposalsService {

  public PendingIdeasList: any[] = []

  constructor(public http: HttpClient) { }

  public loadPendingIdeas() {
    this.http.get<any>(`${apiip}/proposedideas`, { responseType: 'json' })
      .toPromise()
      .then(res => {
        console.log(res);
        this.PendingIdeasList = res
        // console.log(this.PendingIdeasList);

      })
      .catch(err => {
        console.log(err);
      })
  }

  public proposeIdea(idea: any) {
    //   'BlockData': {
    //     "Title": request.json['Title'],
    //     "Description": request.json['Description'],
    //     "Owners": request.json['Owners'],
    //     "Problem": request.json['Problem'],
    //     "Domain": [],
    //     "Technologies_used": request.json['Technologies_used'],
    //     "Viewing_price": request.json['Viewing_price'],
    //     "Ownership_price": request.json['Ownership_price'],
    //     "Pricing_history": request.json['Pricing_history']
    // },
    alert("hello")
    this.http.post(`${nlpapi}/proposeidea`, {
      "Title": idea.Title,
      "Description": idea.Description,
      "Owners": idea.Owners,
      "Problem": idea.Problem,
      "Domain": idea.Domain,
      "Technologies_used": idea.Technologies_used,
      "Viewing_price": idea.Viewing_price,
      "Ownership_price": idea.Ownership_price,
      "Pricing_history": idea.Pricing_history
    })
      .toPromise()
      .then(res => {
        console.log(res);
        this.loadPendingIdeas()
      })
      .catch(err => {
        console.log(err);
        
      })
  }

  public approveIdea(idea: any) {
    this.http.post(`${apiip}/addidea`, idea)
      .toPromise()
      .then(res => {
        console.log(res);
        this.loadPendingIdeas()
      })
      .catch(err => {
        console.log(err);
      })
  }

  public disapproveIdea(idea: any) {
    this.http.post(`${apiip}/disapproveidea`, idea)
      .toPromise()
      .then(res => {
        console.log(res);
        this.loadPendingIdeas()
      })
      .catch(err => {
        console.log(err);
      })
  }


}
