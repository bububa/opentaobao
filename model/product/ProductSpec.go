package product

// ProductSpec 结构体
type ProductSpec struct {
	// 产品规格ID。
	SpecId int64 `json:"spec_id,omitempty" xml:"spec_id,omitempty"`
	// 1:表示可以使用的数据，也就是审核通过的。<br/>3：表示等待小二审核的产品规格，这个数据暂时还不能使用，要等待审核通过后，才能使用。
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 产品品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 产品ID。
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品规格的销售属性组合。格式为：pid:vid;pid:vid。其中：pid是销售属性id，vid是销售属性值id。如果该类目品牌下面没有销售属性，可以不用填写。销售属性通过tmall.brandcat.salespro.get接口获取。
	SpecProps string `json:"spec_props,omitempty" xml:"spec_props,omitempty"`
	// 用户输入的属性值存放位置，例如可输入的销售属性，当用户获取pid vid后，应该先从spec_props_alias中获取，然后通过类目属性获取，获取不到，可以通过这个字段获取。
	CustomePropsName string `json:"custome_props_name,omitempty" xml:"custome_props_name,omitempty"`
	// 销售属性值别名。格式为：pid1:vid1:别名1;pid2:vid2:别名2。其中：pid是销售属性id，vid是销售属性值id。别名长度不可以超过30个字符。目前只有颜色销售属性支持别名。
	SpecPropsAlias string `json:"spec_props_alias,omitempty" xml:"spec_props_alias,omitempty"`
	// 产品的主图片地址。绝对地址，格式：http://host/image_path。
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 认证图片列表
	CertifiedPics []CertPicInfo `json:"certified_pics,omitempty" xml:"certified_pics>cert_pic_info,omitempty"`
	// 产品规格条形码，支持EAN-13格式。
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 产品货号
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 上市时间
	MarketTime string `json:"market_time,omitempty" xml:"market_time,omitempty"`
}
