import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { ShortenedComponent } from './shortened/shortened.component';
import { ProfileComponent } from './profile/profile.component';
import { InvalidComponent } from './invalid/invalid.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'shortened', component: ShortenedComponent },
  { path: 'profile', component: ProfileComponent },
  { path: 'invalid', component: InvalidComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }