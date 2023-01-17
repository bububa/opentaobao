package ott

// MemberDto 结构体
type MemberDto struct {
	// characterName
	CharacterName string `json:"character_name,omitempty" xml:"character_name,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// personId
	PersonId string `json:"person_id,omitempty" xml:"person_id,omitempty"`
	// role
	Role string `json:"role,omitempty" xml:"role,omitempty"`
}
