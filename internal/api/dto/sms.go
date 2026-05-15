package dto

import "errors"

type SendSmsRequest struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

func (r SendSmsRequest) Validate() error {
	if r.Recipient == "" {
		return errors.New("recipient is required")
	}

	if r.Message == "" {
		return errors.New("message is required")
	}

	return nil
}

type SendSmsResponse struct {
	JobID  string `json:"job_id"`
	Status string `json:"status"`
}
