package tmallservice

// WorkerSaveForTopReqDto 结构体
type WorkerSaveForTopReqDto struct {
	// 身份证号码
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 用户地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 用户地址编码
	AddressId int64 `json:"address_id,omitempty" xml:"address_id,omitempty"`
	// 真实姓名
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 工人技能参数
	WorkerServiceAbility *WorkerServiceAbility `json:"worker_service_ability,omitempty" xml:"worker_service_ability,omitempty"`
	// 加入的网点参数
	JoinedStore *JoinedStore `json:"joined_store,omitempty" xml:"joined_store,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}
