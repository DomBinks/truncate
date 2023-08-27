import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  user: string = '';
  rows: Array<Array<string>> = [];

  ngOnInit() {
    interface row {
      original: string;
      short: string;
    }

    this.http.post<any>('/get-profile', {}).subscribe({
      next: response => {
        for (const row of response) {
          this.rows.push([row[0], row[1]]);
        }

        console.log(response)
        console.log(this.rows[0][0]);
        console.log(this.rows[0][1]);
      },
      error: err => {
        console.log("Error: " + err);
      }
    });
  }

  constructor(private http:HttpClient) {}
}
