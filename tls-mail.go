package main

import (
    "crypto/tls"
    "io"
    "net/smtp"
)

const USER = "sean.thorne@gmail.com"
const PASS = "boom"
const MAIL = "email.server.com"
const PORT = "465"

func main() {
    
    To := "test@gmail.com"
    From := USER
    Message := "Testing Testing Testing"
    
    auth := smtp.PlainAuth(USER, USER, PASS, MAIL)
    
    conn, err := tls.Dial("tcp", MAIL + ":" + PORT, &tls.Config{ServerName: MAIL, ClientAuth: tls.NoClientCert})
    
    if err != nil {
        return err
    }
    
    c, e := smtp.NewClient(conn, MAIL)
    if e != nil {
        return e
    }
    
    c.Auth(auth)
    c.Mail(From)
    c.Rcpt(To)
    
    // mail body
    wc, e := c.Data()
    if e != nil {
        return e
    }
    
    defer wc.Close()
    
    _, e = io.WriteString(wc, Message)
    
    if e != nil {
        return e
    }
    
    return nil
}
