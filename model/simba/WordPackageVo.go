package simba

// WordPackageVo 结构体
type WordPackageVo struct {
	// 词包策略信息
	StrategyList []WordPackageStrategyVo `json:"strategy_list,omitempty" xml:"strategy_list>word_package_strategy_vo,omitempty"`
	// 词包名称
	WordPackageName string `json:"word_package_name,omitempty" xml:"word_package_name,omitempty"`
	// 词包基础出价
	BidPrice string `json:"bid_price,omitempty" xml:"bid_price,omitempty"`
	// 计划id,计划已经存在场景必填,eg:添加主体、编辑计划状态等场景
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 词包id
	WordPackageId int64 `json:"word_package_id,omitempty" xml:"word_package_id,omitempty"`
	// 词包类型,0:流量智选,1:好词优选,19:捡漏,20:卖点词包,21:类目优选
	WordPackageType int64 `json:"word_package_type,omitempty" xml:"word_package_type,omitempty"`
	// 词包状态,1:在线,0:下线
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
