package cleaner

import (
	"fmt"
	"strings"

	"github.com/pivotalservices/cf-mgmt/cloudcontroller"
	"github.com/pivotalservices/cf-mgmt/ldap"
	"github.com/pivotalservices/cf-mgmt/uaac"
	"github.com/pivotalservices/cf-mgmt/utils"
	"github.com/xchapter7x/lo"
)

//NewManager -
func NewManager(systemDomain, token, uuacToken string) (mgr Manager) {
	return &DefaultReaperManager{
		UaacMgr:         uaac.NewManager(sysDomain, uaacToken),
		CloudController: cloudcontroller.NewManager(fmt.Sprintf("https://api.%s", sysDomain), token),
		UtilsMgr:        utils.NewDefaultManager(),
		LdapMgr:         ldap.NewManager(),
	}
}

func (m *DefaultReaperManager) CleanEnvironment(printOnly bool) (err error) {
	desiredState := m.buildDesiredState()
	actualState := m.buildDesiredState()
	plan := BuildCleaningPlan(actualState, desiredState)
	plan.RunCleaningPlan()
}

//DefaultUAACManager -
type DefaultReaperManager struct {
	spaceReaper SpaceReaper
}

func (m *DefaultReaperManager) buildActualState() (actualState State, err error) {
	var orgs Organizations
	cfOrgs := m.CloudController.ListOrgs()
	cfSpaces := m.CloudController.ListAllSpaces()

	return actualState
}

func (m *DefaultReaperManager) buildDesiredState() (desiredState State, err error) {
	return desiredState
}

func (m *DefaultReaperManager) buildPlan(actualState, desiredState State) (plan Plan, err error) {
	return
}
