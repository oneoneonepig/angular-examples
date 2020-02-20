import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent{
  ipAddress;
  port;
  username;
  password;

  /*    function checkSqlConnection(ipAddress: string, username: string, password: string) {
    return 'Connecting to ' + ipAddress + ' with username ' + username;
    }*/
}
