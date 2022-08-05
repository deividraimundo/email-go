package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

// PRIMEIRA FORMA DE ENVIAR UM EMAIL

func main() {
	EnviaMail1()
	EnviaMail2()
	EnviaMail3()
}

// Enviando e-mail sem biblioteca
func EnviaMail1() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	from := os.Getenv("MAIL")       // Quem está enviando
	password := os.Getenv("PASSWD") // Senha de que está enviando

	toList := []string{"deivid.raimundo@hotmail.com"} // Quem vai receber

	host := "smtp.aswti.com.br" // Servidor
	port := "587"               // Porta

	msg := "Deu certo 1" // Mensagem

	body := []byte(msg) // Corpo da mensagem

	auth := smtp.PlainAuth("", from, password, host) // Authorization

	err = smtp.SendMail(host+":"+port, auth, from, toList, body) // Enviando e-mail
	if err != nil {
		fmt.Printf("erro ao enviar e-mail: %v", err)
		os.Exit(1)
	}

	fmt.Println("Sucesso ao enviar e-mail!")
}

// Enviando e-mail com biblioteca - VAI PARA SPAM
func EnviaMail2() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", "deivid.raimundo@awsti.com.br") // Quem enviará
	m.SetHeader("To", "deivid.raimundo@hotmail.com")    // Quem receberá
	m.SetHeader("Subject", "Gomail teste")              // Assunto
	m.SetBody("text/plain", "Deu certo 2")              // Mensagem

	d := gomail.NewDialer("smtp.aswti.com.br", 587, os.Getenv("MAIL"), os.Getenv("PASSWD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Sucesso ao enviar e-mail!")
}

// Enviando e-mail através de HTML
func EnviaMail3() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	from := os.Getenv("MAIL")       // Quem vai enviar
	password := os.Getenv("PASSWD") // Senha de quem está enviando

	to := []string{ // Quem vai receber
		"deivid.raimundo@hotmail.com",
	}

	smtpHost := "smtp.aswti.com.br" // Servidor
	smtpPort := "587"               // Porta

	auth := smtp.PlainAuth("", from, password, smtpHost) // Authentication

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Teste \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Deivid",
		Message: "Deu certo 3",
	})

	// Enviando e-mail
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Printf("erro ao enviar e-mail: %v", err)
		os.Exit(1)
	}

	fmt.Println("Sucesso ao enviar e-mail!")
}
