package user

import (
	"github.com/pivotalservices/cf-mgmt/cloudcontroller"
	"github.com/pivotalservices/cf-mgmt/ldap"
	"github.com/pivotalservices/cf-mgmt/uaac"
	// "github.com/pivotalservices/cf-mgmt/utils"
	"github.com/xchapter7x/lo"
)

//NewManager -
func NewManager(sysDomain, token, uaacToken string) (mgr Manager) {
	return &DefaultUserManager{
		UAACMgr:         uaac.NewManager(sysDomain, uaacToken),
		CloudController: cloudcontroller.NewManager(fmt.Sprintf("https://api.%s", sysDomain), token),
		UtilsMgr:        utils.NewDefaultManager(),
		LdapMgr:         ldap.NewManager(),
	}
}

func (m *DefaultOrgManager) AddUserToOrgAndRole(userID, orgGUID, role string) error {
	lo.G.Info("Adding user to groups")
	if err := m.CloudController.AddUserToOrg(userID, orgGUID); err != nil {
		return err
	}
	if err := m.CloudController.AddUserToOrgRole(userID, role, orgGUID); err != nil {
		return err
	}
	return nil
}

func (m *DefaultOrgManager) updateLdapUsers(config *ldap.Config, org *cloudcontroller.Org, role string, uaacUsers map[string]string, users []ldap.User) error {
	for _, user := range users {
		if _, userExists := uaacUsers[strings.ToLower(user.UserID)]; userExists {
			lo.G.Info("User", user.UserID, "already exists")
		} else {
			lo.G.Info("User", user.UserID, "doesn't exist so creating in UAA")
			if err := m.UAACMgr.CreateLdapUser(user.UserID, user.Email, user.UserDN); err != nil {
				lo.G.Error(err)
				return err
			} else {
				uaacUsers[strings.ToLower(user.UserID)] = user.UserID
			}
		}
		if err := m.addUserToOrgAndRole(user.UserID, org.MetaData.GUID, role); err != nil {
			return err
		}
	}
}

func (m *DefaultUserManager) GetLdapUsers(ldapGroupName string, ldapUsersList []string) ([]ldap.User, error) {
	users := []ldap.User{}
	if ldapGroupName != "" {
		if groupUsers, err := m.LdapMgr.GetGroupUsers(config, ldapGroupName); err == nil {
			users = append(users, groupUsers...)
		} else {
			lo.G.Error(err)
			return nil, err
		}
	}
	for _, user := range ldapUsersList {
		if ldapUser, err := m.LdapMgr.GetUser(config, user); err == nil {
			if ldapUser != nil {
				users = append(users, *ldapUser)
			}
		} else {
			lo.G.Error(err)
			return nil, err
		}
	}
	return users, nil
}
