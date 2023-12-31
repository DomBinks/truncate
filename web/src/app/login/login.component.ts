import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  label: string = ''; // Label of the login/logout button
  link: string = ''; // Link the login/logout button should use
  profile: boolean = false; // Whether the profile button should be shown

  // Run when the login/logout button is pressed
  logout(label: string) {
    // If the current function is to logout
    if(label == "Logout") {
      // Delete the authentication cookies
      this.cookieService.delete('auth-session');
    }
  }

  ngOnInit() {
    // Interface for the response to the POST request
    interface loginUI {
      label: string; // Label of the login/logout button
      link: string; // The link the button should use
    }

    // Send a POST request to the server to get the text for the UI
    this.http.get<loginUI>('/get-login').subscribe({
      next: response => {
        // Set the class variables to the values returned
        this.label = response.label;
        this.link = response.link;

        // The profile button should be displayed if the user is logged in
        // i.e. the button's label says logout
        this.profile = response.label == "Logout";
      },
      error: err => {
        console.log("Error: " + err);
      }
    });
  }

  constructor(private http: HttpClient, private cookieService: CookieService) {}
}
