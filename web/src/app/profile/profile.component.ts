import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  // Stores all the user's shortened URLs alongside the original URL
  // Index 0 - original URL, index 1 - shortened URL
  rows: Array<Array<string>> = [];

  ngOnInit() {
    // Interface for the response of the POST request
    interface row {
      original: string;
      short: string;
    }

    // Send a POST request to the server to get the current user's URLs
    this.http.post<any>('/get-urls', {}).subscribe({
      next: response => {
        // For each row in the response
        for (const row of response) {
          // Add the original URL and shortened URL to the rows array
          this.rows.push([row[0], row[1]]);
        }
      },
      error: err => {
        console.log("Error: " + err);
        this.router.navigate(["/"]);
      }
    });
  }

  // When the user selects a shortened URL to delete
  deleteRow(shortened: string) {
    // Send a POST request to the server, specifying the shortened
    // URL of the row to remove
    this.http.post<any>('/delete-row', {"shortened": shortened}).subscribe({
      next: _ => {
        // Refresh the page once the row has been deleted
        window.location.reload();
      },
      error: err => {
        console.log("Error: " + err);
      }
    })
  }

  constructor(private http:HttpClient, private router: Router) {}
}
