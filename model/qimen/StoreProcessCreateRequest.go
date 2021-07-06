package qimen

// StoreProcessCreateRequest 结构体
type StoreProcessCreateRequest struct {
	// 加工商品列表
	Materialitems []MaterialItem `json:"materialitems,omitempty" xml:"materialitems>material_item,omitempty"`
	// 商品列表
	Productitems []ProductItem `json:"productitems,omitempty" xml:"productitems>product_item,omitempty"`
	// 加工单编码
	ProcessOrderCode string `json:"processOrderCode,omitempty" xml:"processOrderCode,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 单据类型(CNJG=仓内加工作业单)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 加工单创建时间(YYYY-MM-DD HH:MM:SS)
	OrderCreateTime string `json:"orderCreateTime,omitempty" xml:"orderCreateTime,omitempty"`
	// 计划加工时间(YYYY-MM-DD HH:MM:SS)
	PlanTime string `json:"planTime,omitempty" xml:"planTime,omitempty"`
	// 加工类型(1:仓内组合加工 2:仓内组合拆分)
	ServiceType string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 成品计划数量
	PlanQty int64 `json:"planQty,omitempty" xml:"planQty,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStoreprocessCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
