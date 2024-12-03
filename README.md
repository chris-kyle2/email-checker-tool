Here’s the updated `README.md` file with your repository URL:

```markdown
# Email Configuration Checker

This is a simple command-line tool written in Go to check the email configuration of domains. It verifies the presence of DNS records like MX, SPF, and DMARC, and outputs the results in a CSV format.

## Features

- Checks if a domain has:
  - **MX Records**: Required for receiving emails.
  - **SPF Records**: Used to prevent email spoofing.
  - **DMARC Records**: Helps protect against phishing and spoofing.
- Outputs results in a CSV format for easy analysis.

## Prerequisites

- [Go](https://go.dev/) installed on your system (version 1.18 or later).

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/chris-kyle2/email-checker-tool.git
   ```
2. Navigate to the project directory:
   ```bash
   cd email-checker-tool
   ```
3. Build the tool:
   ```bash
   go build -o email-checker
   ```

## Usage

1. Run the tool:
   ```bash
   ./email-checker
   ```
2. Enter domain names one per line and press `Enter`:
   ```plaintext
   example.com
   anotherdomain.org
   ```
3. Press `Ctrl+D` (Linux/Mac) or `Ctrl+Z` (Windows) to end input.

### Example Output

The tool outputs results in the following CSV format:

```plaintext
domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord
example.com,true,true,"v=spf1 include:_spf.google.com ~all",true,"v=DMARC1; p=none;"
anotherdomain.org,false,false,"",false,""
```

### Notes

- Ensure you have an active internet connection, as the tool performs live DNS lookups.
- Errors (e.g., if a domain is unreachable) are logged to the console.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add feature description"
   ```
4. Push to your branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- Inspired by the need for quick domain email configuration checks.
- Built using Go's powerful standard library for DNS lookups.
```