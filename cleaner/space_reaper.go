package cleaner

import (
// "fmt"
// "strings"

// "github.com/pivotalservices/cf-mgmt/http"
)

//NewManager -
func NewSpaceManager(systemDomain, uuacToken string) (mgr Manager) {
	return &SpaceManager{}
}

func (m *SpaceManager) CleanSpaces(orgName string) (err error) {
	// get list of all orgs in config repo
	// for each org
	//   get list of all spaces for org in config repo
	//   compare spaces from actual spaces in CF
	//
}

//DefaultUAACManager -
type SpaceManager struct {
	spaceReaper SpaceReaper
}
