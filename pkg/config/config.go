package config

type Mongo struct {
	ConnectionURI    string `default:"mongodb://localhost:27017"`
	BeerCollection   string `default:"beer"`
	ReviewCollection string `default:"review"`
	Database         string `default:"tasmac"`
}

type REST struct {
	Port int `default:"8002"`
}

type HTTP struct {
	REST REST
}

type App struct {
	Mongo Mongo
	HTTP  HTTP
}
