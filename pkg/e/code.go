package e

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	ERROR_REGISTER_FORMAT_FAIL = 20000
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_USER_EXIST = 20006

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003

	ERROR_ADD_BASIC_INFO_FAIL = 40001
	ERROR_GET_BASIC_INFO_FAIL = 40002
)