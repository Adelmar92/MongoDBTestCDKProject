// Retrieves or creates projects in any given Atlas organization
package project


type ProjectApiKey struct {
	Key *string `field:"optional" json:"key" yaml:"key"`
	RoleNames *[]*string `field:"optional" json:"roleNames" yaml:"roleNames"`
}

