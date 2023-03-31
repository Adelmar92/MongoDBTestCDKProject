// Retrieves or creates projects in any given Atlas organization
package project


type ProjectTeam struct {
	RoleNames *[]*string `field:"optional" json:"roleNames" yaml:"roleNames"`
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
}

