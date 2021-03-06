package backends

type Config struct {
	Backend      string
	Prefix       string
	Group        string
	BasicAuth    bool
	ClientCaKeys string
	ClientCert   string
	ClientKey    string
	BackendNodes []string
	Password     string
	Username     string
}
