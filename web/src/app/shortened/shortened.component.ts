import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-shortened',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './shortened.component.html',
  styleUrls: ['./shortened.component.css']
})
export class ShortenedComponent {
  url: string = ''; // Stores the shortened URL returned from the server

  constructor(private route: ActivatedRoute) {
    // Get the query parameters
    this.route.queryParams.subscribe(params => {
      // Get the shortened URL from the query parameters
      this.url = params['shortened'];
    })
  }
}
