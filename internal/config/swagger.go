package config

type Swagger struct {
	SwaggerHost   string `json:"swagger_host"`
	SwaggerScheme string `json:"swagger_scheme"`
	SwaggerPrefix string `json:"swagger_prefix"`
}
