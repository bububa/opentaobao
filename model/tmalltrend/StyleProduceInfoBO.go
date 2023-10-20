package tmalltrend

// StyleProduceInfoBo 结构体
type StyleProduceInfoBo struct {
	// 款式生产信息同步目的，枚举，INSERT(&#34;新增&#34;), UPDATE(&#34;更新&#34;), OFFLINE(&#34;下线&#34;);
	SyncPurpose string `json:"sync_purpose,omitempty" xml:"sync_purpose,omitempty"`
	// 生产价格区间
	ProducePriceRange string `json:"produce_price_range,omitempty" xml:"produce_price_range,omitempty"`
	// 生产商联系方式
	ContactInfo string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	// 款式编号，业务唯一
	StyleSerialNumber string `json:"style_serial_number,omitempty" xml:"style_serial_number,omitempty"`
	// 生产商名称
	ProducerName string `json:"producer_name,omitempty" xml:"producer_name,omitempty"`
	// 工艺源文件地址
	CraftSourceFile string `json:"craft_source_file,omitempty" xml:"craft_source_file,omitempty"`
	// CAD源文件，设计必须
	DesignCadSourceFile string `json:"design_cad_source_file,omitempty" xml:"design_cad_source_file,omitempty"`
	// 生产周期，单位天
	ProduceCycle string `json:"produce_cycle,omitempty" xml:"produce_cycle,omitempty"`
	// 打样价格，单位元
	MakeSamplePrice string `json:"make_sample_price,omitempty" xml:"make_sample_price,omitempty"`
	// 生产商税务编号
	TaxId string `json:"tax_id,omitempty" xml:"tax_id,omitempty"`
	// CAD源文件，生产必须
	ProduceCadSourceFile string `json:"produce_cad_source_file,omitempty" xml:"produce_cad_source_file,omitempty"`
	// 生产Bom清单，json字符串
	Bom string `json:"bom,omitempty" xml:"bom,omitempty"`
	// 最小起订量
	MinimalOder int64 `json:"minimal_oder,omitempty" xml:"minimal_oder,omitempty"`
	// 最大订单量
	MaximumOder int64 `json:"maximum_oder,omitempty" xml:"maximum_oder,omitempty"`
	// 是否支持打样，true--支持，false--不支持
	CanMakeSample bool `json:"can_make_sample,omitempty" xml:"can_make_sample,omitempty"`
}
