import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    HttpClientModule,
  ],
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  url: string = ''; // Stores the URL submitted by the user

  // Called when the shorten button is pressed
  submitURL() {
    const data = {url: this.url}; // Put the URL in a JSON

    // Interface of the response to the POST request
    interface resp {
      shortened: string;
    }

    // Send the JSON to the backend using a POST request
    this.http.post<resp>('/shorten', data).subscribe({
      next: response => {
        // Navigate to the page to display the shortened URL,
        // passing the shortened URL as a query parameter 
        this.router.navigate(['/shortened'], { queryParams: {shortened: response.shortened}});
      },
      error: err => {
        console.log("Error: " + err);

        // If the error was due to submitting an invalid URL
        if(err.status == 400) {
          // Navigate to the page that tells the user the URL was invalid
          this.router.navigate(['/invalid']);
        }
      }
    });
  }

  constructor(private http: HttpClient, private router: Router) {}
}
