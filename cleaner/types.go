package cleaner

import (
	"github.com/pivotalservices/cf-mgmt/cloudcontroller"
	"github.com/pivotalservices/cf-mgmt/ldap"
	"github.com/pivotalservices/cf-mgmt/uaac"
	"github.com/pivotalservices/cf-mgmt/utils"
)

//Manager -
type Manager interface {
	CleanEnvironment() error
}

var ExpectedState State
var ActualState State

type State struct {
	Organizations Organizations
}

type Organizations []Organization
type Spaces []Space
type Actions []Action

type Organization struct {
	Name            string
	Guid            string
	Spaces          Spaces
	BillingManagers Users
	Managers        Users
	Auditors        Users

	LdapBillingManagers string
	LdapManagers        string
	LdapAuditors        string
}

type Space struct {
	Guid string

	Name     string
	AllowSsh bool
	Org      Organization
}

type Users struct {
}

type CleaningPlan struct {
	OrgsToDelete   Organizations
	SpacesToDelete Spaces
	UsersToDelete  Users
}

type Action struct {
	Type   string
	Target interface{}
}
