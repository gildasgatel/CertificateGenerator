package html

import (
	"bytes"
	"certgen/model"
	"certgen/pkg/cert"
	"fmt"
	"strings"
	"text/template"
)

type HtmlCert struct {
}

func New() cert.Service {
	return &HtmlCert{}
}

func (p *HtmlCert) GenCert(data *model.Cert) error {

	tmpl := template.Must(template.New("html").Parse(htmlTemplate))

	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, &data)
	if err != nil {
		return err
	}

	data.Certificat = buffer.Bytes()
	name := strings.ReplaceAll(data.Name, " ", "")
	data.CertName = fmt.Sprintf("cert_%s.html", name)
	return nil
}

var htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Certificat de Réussite</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			padding: 20px;
			text-align: center;
		}
		.certificate {
			border: 2px solid #000;
			padding: 20px;
			margin: 0 auto;
			max-width: 600px;
			background-color: #f9f9f9;
		}
		.name {
			font-size: 24px;
			font-weight: bold;
			margin-bottom: 10px;
		}
		.details {
			font-size: 18px;
			margin-bottom: 15px;
		}
		.date {
			font-style: italic;
		}
	</style>
</head>
<body>
	<div class="certificate">
		<p class="name">{{.Name}}</p>
		<p class="details">{{.Education}}</p>
		<p class="details date">Date: {{.Date}}</p>
	</div>
</body>
</html>`
