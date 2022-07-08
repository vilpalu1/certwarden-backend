package app

import (
	"legocerthub-backend/pkg/output"
	"net/http"
)

type appStatus struct {
	Status          string               `json:"status"`
	Environment     string               `json:"environment"`
	Version         string               `json:"version"`
	AcmeDirectories appStatusDirectories `json:"acme_directories"`
}

type appStatusDirectories struct {
	Production string `json:"prod"`
	Staging    string `json:"staging"`
}

// statusHandler writes some basic info about the status of the Application
func (app *Application) statusHandler(w http.ResponseWriter, r *http.Request) (err error) {

	currentStatus := appStatus{
		Status:  "Available",
		Version: version,
		AcmeDirectories: appStatusDirectories{
			Production: acmeProdUrl,
			Staging:    acmeStagingUrl,
		},
	}

	_, err = output.WriteJSON(w, http.StatusOK, currentStatus, "status")
	if err != nil {
		app.logger.Error(err)
		return output.ErrWriteJsonFailed
	}

	return nil
}

// notFoundHandler is called when there is not a matching route on the router
func (app *Application) notFoundHandler(w http.ResponseWriter, r *http.Request) (err error) {
	_, err = output.WriteErrorJSON(w, output.ErrNotFound)
	if err != nil {
		app.logger.Error(err)
		return output.ErrWriteJsonFailed
	}

	return nil
}
