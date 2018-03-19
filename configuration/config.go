package config

const(
	DB_HOST="tcp(127.0.0.1:3306)"
	DB_NAME="mydb"
	DB_USER="root"
	DB_PASS="root"

// all messages
	INVALID_FORMAT = "Api not defined in URL"
	INVALID_METHOD_MESSAGE = "Invalid method"
	RECORD_NOT_INSERTED = "Record insertion failed"
	RECORD_NOT_UPDATED = "Record updation failed"
	RECORD_NOT_DELETED = "Record deletion failed"
	INVALID_REQUEST_DATA = "Request data is not valid"	

// all status codes
	HTTP_OK = "200"
	HTTP_NOT_FOUND = "404"
	INTERNAL_SERVER_ERROR = "500"
	INVALID_METHOD_CODE = "405"
)
