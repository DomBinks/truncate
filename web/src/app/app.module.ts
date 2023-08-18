import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { ShortenComponent } from './shorten/shorten.component';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    ShortenComponent,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
