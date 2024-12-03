// Package main provides a simple tool to check domain configurations
// for email-related records such as MX, SPF, and DMARC.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// main is the entry point of the application. It reads domains from
// standard input, checks their DNS configurations, and prints the results.
func main() {
	// Create a scanner to read domains from standard input.
	scanner := bufio.NewScanner(os.Stdin)

	// Print the CSV header for the output.
	fmt.Printf("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord\n")

	// Process each domain entered by the user.
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	// Handle errors that may occur while reading from standard input.
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

// checkDomain checks the specified domain for its email configuration,
// including the presence of MX, SPF, and DMARC records.
// It prints the results as a CSV row.
//
// Parameters:
//   - domain: the domain name to check.
func checkDomain(domain string) {
	// Initialize variables to store record presence and their values.
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Look up MX records for the domain.
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Look up TXT records for the domain to find SPF configuration.
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Check each TXT record for an SPF record.
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Look up TXT records for DMARC configuration.
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Check each DMARC record for validity.
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// Print the results as a CSV row.
	fmt.Printf("%v,%v,%v,%v,%v,%v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
