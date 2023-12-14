# CertificatGenerator

CertificatGenerator is a Golang program designed to generate certificates in PDF or HTML format from a CSV file using a Command-Line Interface (CLI).
![This is an image](https://github.com/gildasgatel/CertificateGenerator/blob/master/out/cert_JaneSmith.pdf)


## Features

- **Certificate Generation:** The program takes a CSV file as input and produces certificates in either PDF or HTML format based on the specified parameter.
- **Supported Formats:** Generates certificates in PDF or HTML format based on the specified type.
- **CLI Interface:** Simple usage via a command-line interface to specify the certificate type and CSV file path.

## Usage
To use the program, execute the following command in your terminal:

```bash
./certgen -type <certificate_type> -file /path/to/your/file.csv
```
## Available Options:
- **Certificate Type:** Specify the type of certificate to generate using the **`-type`** option. For example:

  -**type pdf to** generate certificates in PDF format.
  
  -**type html** to generate certificates in HTML format.
- **CSV File Path:** Use the -file option to specify the path to your CSV file containing the data for the certificates.

- Example:

```bash
./certgen -type pdf -file /path/to/your/file.csv
```
## CSV File Structure Example

The CSV file used by CertificatGenerator should follow a specific structure. Here's an example:

```csv
enseignement,name,date
Mathematics,John Doe,2023-12-15
Physics,Jane Smith,2023-12-16
History,Alice Johnson,2023-12-17
