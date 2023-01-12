package contants

type Status int

var (
	StatusOK           Status = 200
	StatusFailOnClient Status = 400
	StatusFailOnServer Status = 500
)
