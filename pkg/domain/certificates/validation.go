package certificates

import (
	"legocerthub-backend/pkg/storage"
	"legocerthub-backend/pkg/validation"
)

// isIdExisting returns an error if not valid. Cert is returned if
// valid
func (service *Service) isIdExisting(id int) (cert Certificate, err error) {
	cert, err = service.storage.GetOneCertById(id, false)
	if err != nil {
		return Certificate{}, err
	}

	return cert, nil
}

// isIdExisting returns an error if not valid. Cert is returned if
// valid
func (service *Service) isIdExistingMatch(idParam int, idPayload *int) (cert Certificate, err error) {
	// basic check
	err = validation.IsIdExistingMatch(idParam, idPayload)
	if err != nil {
		return Certificate{}, err
	}

	// check id exists in storage
	cert, err = service.isIdExisting(idParam)
	if err != nil {
		return Certificate{}, err
	}

	return cert, nil
}

// isNameValid returns an error if not valid, nil if valid
func (service *Service) isNameValid(idPayload *int, namePayload *string) error {
	// basic check
	err := validation.IsNameValid(namePayload)
	if err != nil {
		return err
	}

	// make sure the name isn't already in use in storage
	account, err := service.storage.GetOneCertByName(*namePayload, false)
	if err == storage.ErrNoRecord {
		// no rows means name is not in use
		return nil
	} else if err != nil {
		// any other error, return the error
		return err
	}

	// if the returned account is the account being edited, no error
	if *account.ID == *idPayload {
		return nil
	}

	return validation.ErrNameInUse
}
