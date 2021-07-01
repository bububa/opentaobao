package wdk

// ErpArrivalNoticeDetailDto 结构体
type ErpArrivalNoticeDetailDto struct {
	// 采购单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 库存单位
	InventoryUnit string `json:"inventory_unit,omitempty" xml:"inventory_unit,omitempty"`
	// 规格
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 部门code，该商品所属的部门编码
	DeptCode string `json:"dept_code,omitempty" xml:"dept_code,omitempty"`
	// 1
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 库存单位
	PlanPackageQuantity string `json:"plan_package_quantity,omitempty" xml:"plan_package_quantity,omitempty"`
	// 数量
	Count string `json:"count,omitempty" xml:"count,omitempty"`
	// 生产日期： 1、 基础规则：生产日期不得超过或等于今日； 2、 商品未设置保质期管理的，生产日期为可选项，UMS存储但不校验禁收时限： a) 单据传输中提供了生产日期的，按单据生产日期通过收货； b) 单据传输中未提供生产日期的，不存储生产日期通过收货； 3、 商品已设置保质期管理的，并且已经设置了禁收时限的： a) 单据传输中提供了生产日期的，校验是否符合禁收时限，通过收货，不通过整单不收货； b) 单据传输中未提供生产日期的，按生产日期＝入库时间－禁收时限，通过收货； 4、 商品已设置保质期管理的，但未设置禁收时限的： a) 单据传输中提供了生产日期的，按单据生产日期，通过收货； b) 单据传输中未提供生产日期的，按生产日期＝入库时间－1（天），通过收货。
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 商品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 商品code，盒马系统中的商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 可指定库位，可空（按照需求附值）
	CabinetCode string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
}
