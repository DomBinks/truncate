import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-shorten',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    HttpClientModule,
  ],
  templateUrl: './shorten.component.html',
  styleUrls: ['./shorten.component.css']
})
export class ShortenComponent {
  url:string = '';

  submitURL() {
    console.log(this.url);

    const data = {url: this.url};
    this.http.post('/shorten', data);
    this.router.navigate(['/shortened']);
  }

  constructor(private http: HttpClient, private router: Router) {}
}
