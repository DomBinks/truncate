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
  url: string = ''; // Stores the URL submitted by the user

  submitURL() {
    const data = {url: this.url}; // Put the URL in a JSON

    // Interface of the response to the POST request
    interface resp {
      short: string;
    }

    // Send a POST request with the JSON to the Go backend
    // (Have to subscribe for the POST request to be sent)
    this.http.post<resp>('/shorten', data).subscribe({
      next: response => {
        // Navigate to the page to display the shortened link
        this.router.navigate(['/shortened'], { queryParams: {short: response.short}});
      },
      error: err => {
        console.log("Error: " + err);
      }
    });
  }

  constructor(private http: HttpClient, private router: Router) {}
}
