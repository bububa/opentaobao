package simba

// AdzoneRefVo 结构体
type AdzoneRefVo struct {
	// 操作按钮列表
	OperationList []Integer `json:"operation_list,omitempty" xml:"operation_list>integer,omitempty"`
	// 资源包名字
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 出价方式,custom_bid:手动出价,max_amount:最大化拿量,cost_control:控成本,roi_control:控投产比
	BidType string `json:"bid_type,omitempty" xml:"bid_type,omitempty"`
	// 状态,start:投放,pause:暂停
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 资源包id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 计划id,计划已经存在场景必填,eg:添加主体、编辑计划状态等场景
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 资源位/包 溢价
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 资源位支持的最低溢价
	MinDiscount int64 `json:"min_discount,omitempty" xml:"min_discount,omitempty"`
	// 资源位支持的最高溢价
	MaxDiscount int64 `json:"max_discount,omitempty" xml:"max_discount,omitempty"`
}
