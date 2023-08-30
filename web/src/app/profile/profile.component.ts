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
      },
      error: err => {
        console.log("Error: " + err);
        this.router.navigate(["/"]);
      }
    });
  }

  deleteRow(row: string) {
    this.http.post<any>('/delete-row', {"row": row}).subscribe({
      next: _ => {
        window.location.reload();
      },
      error: err => {
        console.log("Error: " + err);
      }
    })
  }

  constructor(private http:HttpClient, private router: Router) {}
}
