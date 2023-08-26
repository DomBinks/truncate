import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ShortenComponent } from './shorten/shorten.component';
import { LinkComponent } from './link/link.component';
import { ProfileComponent } from './profile/profile.component';

const routes: Routes = [
  { path: '', component: ShortenComponent},
  { path: 'shortened', component: LinkComponent},
  { path: 'profile', component: ProfileComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }