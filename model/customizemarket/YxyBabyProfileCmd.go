package customizemarket

// YxyBabyProfileCmd 结构体
type YxyBabyProfileCmd struct {
	// 出生年月或者预产期
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 性别
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 授权id
	ProfileId int64 `json:"profile_id,omitempty" xml:"profile_id,omitempty"`
	// 宝贝类型
	BabyType int64 `json:"baby_type,omitempty" xml:"baby_type,omitempty"`
}
