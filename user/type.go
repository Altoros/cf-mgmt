package user

import (
// "github.com/pivotalservices/cf-mgmt/cloudcontroller"
// "github.com/pivotalservices/cf-mgmt/ldap"
// "github.com/pivotalservices/cf-mgmt/uaac"
// "github.com/pivotalservices/cf-mgmt/utils"
)

//Manager -
type Manager interface {
	// FindOrg(orgName string) (org *cloudcontroller.Org, err error)
	// CreateOrgs(configFile string) (err error)
	// UpdateOrgUsers(configDir, ldapBindPassword string) (err error)
	// CreateQuotas(configDir string) (err error)
	// GetOrgGUID(orgName string) (orgGUID string, err error)
}

//DefaultManager -
type DefaultUserManager struct {
	CloudController cloudcontroller.Manager
	UAACMgr         uaac.Manager
	UtilsMgr        utils.Manager
	LdapMgr         ldap.Manager
	LdapConfig      *ldap.Config
}
