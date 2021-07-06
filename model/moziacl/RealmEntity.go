package moziacl

// RealmEntity 结构体
type RealmEntity struct {
	// 租户名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 租户描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 租户id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
