import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FourOhfourComponent } from './four-ohfour.component';

describe('FourOhfourComponent', () => {
  let component: FourOhfourComponent;
  let fixture: ComponentFixture<FourOhfourComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FourOhfourComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FourOhfourComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
