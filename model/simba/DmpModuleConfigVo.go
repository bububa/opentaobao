package simba

// DmpModuleConfigVo 结构体
type DmpModuleConfigVo struct {
	// dmp数据来源选项
	DmpBaseOptionalSelectAppList []DmpBaseOptionalSelectVo `json:"dmp_base_optional_select_app_list,omitempty" xml:"dmp_base_optional_select_app_list>dmp_base_optional_select_vo,omitempty"`
	// 投放渠道id，筛在达摩盘平台上创建且同步到业务线的人群
	DeliverAppId int64 `json:"deliver_app_id,omitempty" xml:"deliver_app_id,omitempty"`
}
