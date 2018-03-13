package config

const(
	DB_HOST="tcp(127.0.0.1:3306)"
	DB_NAME="mydb"
	DB_USER="root"
	DB_PASS="root"

// all messages
	INVALID_FORMAT = "Invalid format of Request"
	INVALID_METHOD_MESSAGE = "Invalid method"

// all status codes
	HTTP_OK = "200"
	HTTP_NOT_FOUND = "404"
	INTERNAL_SERVER_ERROR = "500"
	INVALID_METHOD_CODE = "405"


)
