import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-link',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './link.component.html',
  styleUrls: ['./link.component.css']
})
export class LinkComponent {
  short: string = ''; // Stores generated number provided by the query parameters

  constructor(private route: ActivatedRoute) {
    // Get the query parameters
    this.route.queryParams.subscribe(params => {
      // Get the short query parameter
      this.short = params['short'];
    })
  }
}
