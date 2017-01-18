package cloudcontroller

import "github.com/pivotalservices/cf-mgmt/http"

type Manager interface {
	CreateSpace(spaceName, orgGUID string) error
	ListSpaces(orgGUID string) ([]Space, error)
	AddUserToSpaceRole(userName, role, spaceGUID string) error
	UpdateSpaceSSH(sshAllowed bool, spaceGUID string) error

	AssignSecurityGroupToSpace(spaceGUID, sgGUID string) error
	ListSecurityGroups() (map[string]string, error)
	CreateSecurityGroup(sgName, contents string) (string, error)
	UpdateSecurityGroup(sgGUID, sgName, contents string) error

	CreateSpaceQuota(orgGUID, quotaName string,
		memoryLimit, instanceMemoryLimit, totalRoutes, totalServices int,
		paidServicePlansAllowed bool) (string, error)
	UpdateSpaceQuota(orgGUID, quotaGUID, quotaName string,
		memoryLimit, instanceMemoryLimit, totalRoutes, totalServices int,
		paidServicePlansAllowed bool) error
	ListSpaceQuotas(orgGUID string) (map[string]string, error)
	AssignQuotaToSpace(spaceGUID, quotaGUID string) error

	CreateOrg(orgName string) error
	ListOrgs() ([]*Org, error)
	AddUserToOrgRole(userName, role, orgGUID string) error
	RemoveUserFromOrgRole(orgGUID, userId, role string) error
	ListUsersWithOrgRole(orgGUID, role string) ([]string, error)
	AddUserToOrg(userName, orgGUID string) error
	// ListOrgUsers(role, orgGUID string) ([]string, error)

	ListQuotas() (quotas map[string]string, err error)
	CreateQuota(quotaName string,
		memoryLimit, instanceMemoryLimit, totalRoutes, totalServices int,
		paidServicePlansAllowed bool) (string, error)
	UpdateQuota(quotaGUID, quotaName string,
		memoryLimit, instanceMemoryLimit, totalRoutes, totalServices int,
		paidServicePlansAllowed bool) error

	AssignQuotaToOrg(orgGUID, quotaGUID string) error
}

type DefaultManager struct {
	Host  string
	Token string
	HTTP  http.Manager
}

//SpaceResources -
type SpaceResources struct {
	Spaces []Space `json:"resources"`
}

//SpaceResources -
type UserResources struct {
	Resources []User `json:"resources"`
}

type User struct {
	MetaData UserMetaData `json:"metadata"`
	Entity   UserEntity   `json:"entity"`
}

//SpaceMetaData -
type UserMetaData struct {
	GUID string `json:"guid"`
}

//SpaceMetaData -
type UserEntity struct {
	Username string
}

type Space struct {
	MetaData SpaceMetaData `json:"metadata"`
	Entity   SpaceEntity   `json:"entity"`
}

//SpaceMetaData -
type SpaceMetaData struct {
	GUID string `json:"guid"`
}

//SpaceEntity -
type SpaceEntity struct {
	Name     string `json:"name"`
	AllowSSH bool   `json:"allow_ssh"`
	OrgGUID  string `json:"organization_guid"`
}

//Orgs -
type Orgs struct {
	Orgs []*Org `json:"resources"`
}

//Org -
type Org struct {
	Entity   OrgEntity   `json:"entity"`
	MetaData OrgMetaData `json:"metadata"`
}

//OrgEntity -
type OrgEntity struct {
	Name string `json:"name"`
}

//OrgMetaData -
type OrgMetaData struct {
	GUID string `json:"guid"`
}

//SecurityGroupResources -
type SecurityGroupResources struct {
	SecurityGroups []SecurityGroup `json:"resources"`
}

type SecurityGroup struct {
	MetaData SecurityGroupMetaData `json:"metadata"`
	Entity   SecurityGroupEntity   `json:"entity"`
}

//SecurityGroupMetaData -
type SecurityGroupMetaData struct {
	GUID string `json:"guid"`
}

//SecurityGroupEntity -
type SecurityGroupEntity struct {
	Name string `json:"name"`
}

//Quotas -
type Quotas struct {
	Quotas []Quota `json:"resources"`
}

type Quota struct {
	MetaData QuotaMetaData `json:"metadata"`
	Entity   QuotaEntity   `json:"entity"`
}

//QuotaMetaData -
type QuotaMetaData struct {
	GUID string `json:"guid"`
}

//QuotaEntity -
type QuotaEntity struct {
	Name string `json:"name"`
}
