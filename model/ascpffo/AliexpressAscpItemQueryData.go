package ascpffo

// AliexpressAscpItemQueryData 结构体
type AliexpressAscpItemQueryData struct {
	// 扩展字段
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
	// 货品声明价值
	CustomsUnitPrice string `json:"customs_unit_price,omitempty" xml:"customs_unit_price,omitempty"`
	// 皮重
	TWeight string `json:"t_weight,omitempty" xml:"t_weight,omitempty"`
	// 净重
	NWeight string `json:"n_weight,omitempty" xml:"n_weight,omitempty"`
	// 毛重
	GWeight string `json:"g_weight,omitempty" xml:"g_weight,omitempty"`
	// 包装材料
	PackageMaterial string `json:"package_material,omitempty" xml:"package_material,omitempty"`
	// 颜色
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// 货品条码
	WhcBarCode string `json:"whc_bar_code,omitempty" xml:"whc_bar_code,omitempty"`
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 货品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 重
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 高
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 宽
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 长度
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 类目Id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 品牌Id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 货品Id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
