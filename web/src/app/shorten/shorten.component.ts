import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-shorten',
  standalone: true,
  imports: [CommonModule],
  template: `
    <p>
      shorten works!
    </p>
  `,
  styleUrls: ['./shorten.component.css']
})
export class ShortenComponent {

}
