package cert

import "certgen/model"

type Service interface {
	GenCert(*model.Cert) error
}
