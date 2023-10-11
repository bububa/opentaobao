package ascp

// ExtendMessage 结构体
type ExtendMessage struct {
	// 地址黑名单原因枚举，仅允许传一种： NEED_SWIPE_CARD  有门禁（刷卡、刷脸、梯控等） SECURITY_REFUSE 保安/门卫不允许进入 UNIT 特殊地址-部队 FACTORY 特殊地址-厂区 URBAN_VILLAGE 特殊地址-城中村 VILLAGE 特殊地址-村镇/村庄 TRANSIT_ WAREHOUSE 特殊地址-转运仓 COLLECTION_POINT 特殊地址-代收点 SCHOOL 特殊地址-学校 OTHER  其他原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 详细原因 条件必填，reason=OTHER时，需要填写eg：都是放到门卫处
	DetailReason string `json:"detail_reason,omitempty" xml:"detail_reason,omitempty"`
}
