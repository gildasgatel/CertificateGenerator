package pdf

import (
	"bytes"
	"certgen/model"
	"certgen/pkg/cert"
	"fmt"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

const (
	IMG_BACKGROUND = "./public/asset/background.png"
	IMG_CERT       = "./public/asset/cert.png"
)

type PdfCert struct {
}

func New() cert.Service {
	return &PdfCert{}
}

func (p *PdfCert) GenCert(data *model.Cert) error {
	name := strings.ReplaceAll(data.Name, " ", "")
	data.CertName = fmt.Sprintf("cert_%s.pdf", name)
	pdf := gofpdf.New("L", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 24)
	w, h := pdf.GetPageSize()
	pdf.Image(IMG_BACKGROUND, 0, 0, w, h, false, "", 0, "")

	// write name
	pdf.Text(100, 75, data.Name)
	// write education
	pdf.Text(140, 125, data.Education)
	// write date
	pdf.SetFontSize(14)
	pdf.Text(50, 185, data.Date)

	// add certificate
	pdf.Image(IMG_CERT, 180, 110, 100, 100, false, "", 0, "")

	// Enregistrer le document PDF dans un fichier
	var buffer bytes.Buffer
	err := pdf.Output(&buffer)
	if err != nil {
		return err
	}
	data.Certificat = buffer.Bytes()
	return nil
}
