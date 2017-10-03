package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	hostspath, isSet := os.LookupEnv("CERTTOWERHOSTS")
	if !isSet {
		homedir, _ := os.LookupEnv("HOME")
		hostspath = path.Join(homedir, ".certificate-tower-hosts")
	}
	f, err := os.Open(hostspath)
	if err != nil {
		log.Fatalf("Could not open hosts file: %v", err)
	}
	defer f.Close()

	var expiredHosts int
	r := bufio.NewReader(f)
	for {
		host, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		host = strings.TrimSpace(host)
		if len(host) == 0 {
			continue
		}
		expiryDate, err := expirationDate(host)
		if x509Err, ok := err.(x509.CertificateInvalidError); ok && x509Err.Reason == x509.Expired {
			fmt.Printf("%s has already expired | color=red \n", host)
			expiredHosts++
			continue
		}
		if err != nil {
			log.Fatalf("encountered error during checking of '%s': %v", host, err)
		}

		if time.Now().Add(14 * 24 * time.Hour).After(expiryDate) {
			fmt.Printf("%s expires on %s | color=orange \n", host, expiryDate.Format("2006-01-02"))
			expiredHosts++
		}
	}
	if expiredHosts == 0 {
		fmt.Println("ok")
	}
}

func expirationDate(domain string) (time.Time, error) {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{})
	if err != nil {
		return time.Now(), err
	}
	return conn.ConnectionState().VerifiedChains[0][0].NotAfter, nil
}
