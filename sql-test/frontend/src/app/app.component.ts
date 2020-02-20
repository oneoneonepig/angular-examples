import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.css']
})
export class AppComponent{
	data = {inputIpAddress: '', inputUsername: '', inputPassword: '', inputTimeout: 3}
	testResult;

	constructor(
		private http: HttpClient
	){}
	
	ngOnInit(){
		
	}

	connTest(data: any) {
		this.http.post('/check', data)
			.subscribe(
				(response) => { console.log(response); this.testResult = response; }
			);
	}
}
