import (

	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"crypto/tls"

)

func gomail(){

    from := mail.Address{"", "from@yahoo.co.jp"}
    to   := mail.Address{"", "to@icloud.com"}
    subj := "件名を日本語化"
    body := "日本語でのメール送信テスト\n 送信できていればOK\n ２"

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // Connect to the SMTP Server
    servername := "smtp.mail.yahoo.co.jp:465"

    host, _, _ := net.SplitHostPort(servername)

    auth := smtp.PlainAuth("","bitcoin_sql", "994545is", host)

    // TLS config
    tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    }

    // Here is the key, you need to call tls.Dial instead of smtp.Dial
    // for smtp servers running on 465 that require an ssl connection
    // from the very beginning (no starttls)
    conn, err := tls.Dial("tcp", servername, tlsconfig)
    if err != nil {
        log.Panic(err)
    }

    c, err := smtp.NewClient(conn, host)
    if err != nil {
        log.Panic(err)
    }

    // Auth
    if err = c.Auth(auth); err != nil {
        log.Panic(err)
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        log.Panic(err)
    }

    if err = c.Rcpt(to.Address); err != nil {
        log.Panic(err)
    }

    // Data
    w, err := c.Data()
    if err != nil {
        log.Panic(err)
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        log.Panic(err)
    }

    err = w.Close()
    if err != nil {
        log.Panic(err)
    }

    c.Quit()

}
