package models

type Config struct {
	dialect    string
	credential string
}

var dbConfig = Config{
	dialect:    "postgres",
	credential: "user=seke1412 dbname=testdb password=seke1412 sslmode=disable",
}
