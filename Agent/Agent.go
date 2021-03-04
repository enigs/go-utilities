package Agent

import (
	"net/http"
	"strings"

	"github.com/tomasen/realip"
	"github.com/varstr/uaparser"

	"github.com/enigs/go-utilities/StringUtils"
)

// Set Agent struct
type Agent struct {
	Authorization           string
	Browser                 string
	BrowserVersion          string
	Ip                      string
	Os                      string
	OsVersion               string
}

// Creates a constructor for agent
func NewAgent() *Agent {
	// Set agent struct
	var a Agent

	// Return newly generated struct
	return &a
}

// Get headers
func (agent *Agent) GetHeaders(r *http.Request) error {
	// Retrieve authorization
	authorization := strings.TrimSpace(r.Header.Get("Authorization"))
	bearer := strings.Split(authorization, " ")

	// Extract token from bearer
	if len(bearer) == 2 && bearer[0] == "Bearer" && !StringUtils.IsEmpty(bearer[1]) {
		agent.Authorization = bearer[1]
	}

	// Retrieve ip
	agent.Ip = realip.FromRequest(r)

	// Check if accessed via browser
	ua := uaparser.Parse(r.UserAgent())

	if ua.Browser != nil {
		agent.Browser = ua.Browser.Name
		agent.BrowserVersion = ua.Browser.Version
	}

	if ua.OS != nil {
		agent.Os = ua.OS.Name
		agent.OsVersion = ua.OS.Version
	}

	return nil
}