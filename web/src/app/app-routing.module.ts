import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ShortenComponent } from './shorten/shorten.component';
import { LinkComponent } from './link/link.component';

const routes: Routes = [
  { path: '', component: ShortenComponent},
  { path: 'shortened', component: LinkComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }