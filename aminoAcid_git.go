package minimin

import (
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type GitParams struct {
	URL       string `yaml:"url" json:"url"`
	User      string `yaml:"user;omitempty" json:"user;omitempty"`
	Password  string `yaml:"password;omitempty" json:"password;omitempty"`
	SecretKey string `yaml:"secret_key;omitempty" json:"secret_key;omitempty"`
	Alias     string `yaml:"alias;omitempty" json:"alias;omitempty"`
	Cell      string `yaml:"-" json:"-"`
}

func GitPull(obj GitParams) (err error) {
	var opt git.CloneOptions
	opt.URL = obj.URL
	var space string
	if obj.Alias == "" {
		var pl = strings.Split(obj.Alias, "/")
		obj.Alias = strings.TrimRight(pl[len(pl)-1], ".git")
	}
	if obj.User != "" && obj.Password != "" {
		opt.Auth = &http.BasicAuth{
			Username: obj.User,
			Password: obj.Password,
		}
	} else if obj.User != "" && obj.SecretKey != "" {
		opt.Auth = &http.BasicAuth{
			Username: obj.User,
			Password: obj.SecretKey,
		}
	}
	space = obj.Cell + "/" + obj.Alias
	if err = os.Mkdir(space, 0777); err != nil {
		return
	}
	_, err = git.PlainClone(space, false, &opt)
	return
}
