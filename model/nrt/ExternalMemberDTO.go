package nrt

// ExternalMemberDto 结构体
type ExternalMemberDto struct {
	// 业态编号
	BusiTypeId string `json:"busi_type_id,omitempty" xml:"busi_type_id,omitempty"`
	// 淘宝昵称 关注过品牌号的会员，有淘宝昵称
	TaoId string `json:"tao_id,omitempty" xml:"tao_id,omitempty"`
	// 手机号
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 地址	地址包含省、市、区、街道、详细地址，仅发最新的一条含有地址的联系信息
	NewAddr string `json:"new_addr,omitempty" xml:"new_addr,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 会员等级
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 会员ID
	MemberId string `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// 操作类型 基础类型 new?、update、delete、query
	OpType string `json:"op_type,omitempty" xml:"op_type,omitempty"`
}
