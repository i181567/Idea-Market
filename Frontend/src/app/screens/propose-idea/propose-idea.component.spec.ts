import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProposeIdeaComponent } from './propose-idea.component';

describe('ProposeIdeaComponent', () => {
  let component: ProposeIdeaComponent;
  let fixture: ComponentFixture<ProposeIdeaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProposeIdeaComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProposeIdeaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
