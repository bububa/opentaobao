package logistic

// JzTopDto 结构体
type JzTopDto struct {
	// 快递公司列表
	Expresses []Expresses `json:"expresses,omitempty" xml:"expresses>expresses,omitempty"`
	// 物流公司列表
	LgCps []Lgcps `json:"lg_cps,omitempty" xml:"lg_cps>lgcps,omitempty"`
	// 安装公司列表
	InsTps []Instps `json:"ins_tps,omitempty" xml:"ins_tps>instps,omitempty"`
	// 商品对应的服务信息
	GoodsRelations string `json:"goods_relations,omitempty" xml:"goods_relations,omitempty"`
	// 是否支持快递
	SupportDelivery bool `json:"support_delivery,omitempty" xml:"support_delivery,omitempty"`
	// 是否支持安装
	SupportInstall bool `json:"support_install,omitempty" xml:"support_install,omitempty"`
	// 是否支持修改安装地址
	SuppModifyInsAdd bool `json:"supp_modify_ins_add,omitempty" xml:"supp_modify_ins_add,omitempty"`
}
