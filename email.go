package main

import (
	"bytes"
	"fmt"
	"html/template"

	gomail "gopkg.in/gomail.v2"
)

func EmailAction(data bookingData, toMail string) {

	from := "harish@infonixweblab.com"
	to := toMail
	subject := "Testing Schedular"
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("subject", subject)
	fmt.Println("!!!!!!", m)

	templateAddr := "F:\\Infonix\\reminder.html"

	m.SetBody("text/html", getTemplate(templateAddr))
	fmt.Println(">>>>>", m)
	d := gomail.NewDialer("smtp.zoho.com", 587, "harish@infonixweblab.com", "harish123")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("Connetion successfull!!")
	fmt.Println("mail sent")

}

func getTemplate(file string) string {
	//function to make body of the email (template + Data)
	t, _ := template.ParseFiles(file)
	buf := new(bytes.Buffer)
	data := "../html/userdetail.html"
	fmt.Println("??????", t)
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}
	return buf.String()

}
