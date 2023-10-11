package cainiaolocker

// PackageInfoDto 结构体
type PackageInfoDto struct {
	// 商品信息,数量限制为100
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 大件快运中的包装方式描述
	PackagingDescription string `json:"packaging_description,omitempty" xml:"packaging_description,omitempty"`
	// 包裹id，用于拆合单场景（只能传入数字、字母和下划线；批量请求时值不得重复，大小写敏感，即123A,123a 不可当做不同ID，否则存在一定可能取号失败）
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 大件快运中的货品描述，比如服装，家具。 顺丰取号必须传此参数
	GoodsDescription string `json:"goods_description,omitempty" xml:"goods_description,omitempty"`
	// 物品价值，单位元
	GoodValue string `json:"good_value,omitempty" xml:"good_value,omitempty"`
	// 包裹长，单位厘米
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 重量,单位 g
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 子母件模式中的总包裹数/总件数，用于打印当前包裹处于总件数的位置比如5-2，可以表示总包裹数为5，当前为第2个包裹，只有快运公司需要传入，其他的可以不用传入
	TotalPackagesCount int64 `json:"total_packages_count,omitempty" xml:"total_packages_count,omitempty"`
	// 体积, 单位 ml
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 包裹宽，单位厘米
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 包裹高，单位厘米
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
