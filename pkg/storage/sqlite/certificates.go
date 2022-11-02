package sqlite

import (
	"legocerthub-backend/pkg/challenges"
	"legocerthub-backend/pkg/domain/certificates"
)

// certificateDb is a single certificate, as database table fields
// corresponds to certificates.Certificate
type certificateDb struct {
	id                   int
	name                 string
	description          string
	certificateKeyDb     keyDb
	certificateAccountDb accountDb
	subject              string
	subjectAltNames      commaJoinedStrings
	challengeMethodValue string
	organization         string
	organizationalUnit   string
	country              string
	state                string
	city                 string
	createdAt            int
	updatedAt            int
	apiKey               string
	apiKeyViaUrl         bool
}

func (cert certificateDb) toCertificate() certificates.Certificate {
	return certificates.Certificate{
		ID:                 cert.id,
		Name:               cert.name,
		Description:        cert.description,
		CertificateKey:     cert.certificateKeyDb.toKey(),
		CertificateAccount: cert.certificateAccountDb.toAccount(),
		Subject:            cert.subject,
		SubjectAltNames:    cert.subjectAltNames.toSlice(),
		ChallengeMethod:    challenges.MethodByValue(cert.challengeMethodValue),
		Organization:       cert.organization,
		OrganizationalUnit: cert.organizationalUnit,
		Country:            cert.country,
		State:              cert.state,
		City:               cert.city,
		CreatedAt:          cert.createdAt,
		UpdatedAt:          cert.updatedAt,
		ApiKey:             cert.apiKey,
		ApiKeyViaUrl:       cert.apiKeyViaUrl,
	}
}
