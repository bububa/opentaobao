package vaccin

// IsvVcAvailableRemindResponse 结构体
type IsvVcAvailableRemindResponse struct {
	// 登记单主键id。total如果大于200，则只显示200个id
	OuterRegisterIdList []string `json:"outer_register_id_list,omitempty" xml:"outer_register_id_list>string,omitempty"`
	// 登记单总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
