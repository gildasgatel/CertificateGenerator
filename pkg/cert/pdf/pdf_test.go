package pdf

import (
	"certgen/model"
	"testing"
)

func TestGenCert(t *testing.T) {
	mockData := &model.Cert{
		Name:      "John Doe",
		Education: "Example Education",
		Date:      "2023-12-31",
	}

	pdfCert := &PdfCert{}

	err := pdfCert.GenCert(mockData)
	if err != nil {
		t.Errorf("Error generating certificate: %v", err)
	}

	if len(mockData.Certificat) == 0 {
		t.Error("Certificate not generated")
	}
}
