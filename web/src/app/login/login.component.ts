import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  func: string = '';
  link: string = '';
  login: boolean = false;

  ngOnInit() {
    interface loginUI {
      func: string;
      link: string;
    }

    this.http.get<loginUI>('/loginUI').subscribe({
      next: response => {
        this.func = response.func;
        this.link = response.link;
        this.login = response.func == "Logout";
      },
      error: err => {
        console.log("Error: " + err);
      }
    });
  }

  constructor(private http: HttpClient) {}
}