package toggly

const API_HOST = "https://api.toggly.co"

var inited = false

type SDK struct {
	Host    string
	UserKey string
	EnvKey  string
}

func (sdk *SDK) Init(user string, env string) {
	sdk.Host = API_HOST
	sdk.UserKey = user
	sdk.EnvKey = env
	if inited {
		return
	}
}

func (sdk *SDK) IsFlagAllowed(key string, instance string) bool {
	return true
}

func (sdk *SDK) GetFlagValueAsInteger(key string, instance string) int {
	return 0
}

func (sdk *SDK) GetFlagValueAsString(key string, instance string) string {
	return ""
}

