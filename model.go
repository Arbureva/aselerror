package aselerror

var (
	NotFound      = New("not found", WithCode(404), WithLevel(1))
	Internal      = New("internal error", WithCode(500), WithLevel(1))
	IllegalParame = New("illegal parame", WithCode(400), WithLevel(1))
)
