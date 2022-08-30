package lstpos

// GoodsDto 结构体
type GoodsDto struct {
	// 规格
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 简称
	ShotTitle string `json:"shot_title,omitempty" xml:"shot_title,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品条码,唯一
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 快捷码(非必须)
	ShortcutCode string `json:"shortcut_code,omitempty" xml:"shortcut_code,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 商品状态 售卖中:sale 不可售卖：not sell
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 计价方式  称重商品: weigh 非称重: non weigh
	PricingMode string `json:"pricing_mode,omitempty" xml:"pricing_mode,omitempty"`
	// 自营叶子类目自定义编号(ISV子定义类目Id)
	IsvCategoryId string `json:"isv_category_id,omitempty" xml:"isv_category_id,omitempty"`
	// 类目Id  1,洗涤用品; 2,家庭日用品; 3,小食品; 4,粮油调味; 5,酒水饮料; 6,速食; 7,生鲜; 8,服务性商品; 9,增值服务; 10,非商品收入;98,无码; 99,其他
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// ISV商品Id(保存商品入参用到)
	IsvGoodsId string `json:"isv_goods_id,omitempty" xml:"isv_goods_id,omitempty"`
	// 零售价，单位为分
	ReservePrice int64 `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 修改时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 进货价 ，单位为分
	OriginalPrice int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
}
