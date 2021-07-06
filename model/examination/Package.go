package examination

// Package 结构体
type Package struct {
	// 体检项列表
	ItemList []Item `json:"item_list,omitempty" xml:"item_list>item,omitempty"`
	// 入职体检,父母体检,女性体检,男性体检,儿童体检,婚前孕前,肿瘤筛查,深度体检,核酸检测,专项检查,白领体检,特价体检,证件体检,皮肤检测,健康证；多个逗号分隔
	Labels []string `json:"labels,omitempty" xml:"labels>string,omitempty"`
	// 三高检测,肿瘤筛查,前列腺检查,宫颈癌筛查,妇科检查,乳腺检查,甲状腺检查,心脑血管检查,胸肺部检查,肝胆检查,专项检查,基因检测,脱发检测,毛囊检测,毛发检测,胃部疾病筛查,肠道疾病筛查,脊椎检测；多个逗号分隔
	FeatureItem []string `json:"feature_item,omitempty" xml:"feature_item>string,omitempty"`
	// 服务类型 ：ONSITE_SERVICE（到店检测）, DOOR_TO_DOOR_SERVICE（上门检测） ; 多种服务类型用逗号分隔
	ServiceTypes []string `json:"service_types,omitempty" xml:"service_types>string,omitempty"`
	// 性别（0：男，1：女，2：通用）
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 售卖价（单位分）（数值需要大于100，且大于settlement_price字段的数值），精确到分，民营商品必需要，并且一个商品组下面的所有套餐售卖价必须一样
	ContractPrice string `json:"contract_price,omitempty" xml:"contract_price,omitempty"`
	// 结算价（单位分）（数值需要大于100），精确到分
	SettlementPrice string `json:"settlement_price,omitempty" xml:"settlement_price,omitempty"`
	// 套餐代码，机构保证全局唯一
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 套餐名称（必填！！！）
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 套餐详情
	PackageDetail string `json:"package_detail,omitempty" xml:"package_detail,omitempty"`
	// 婚否（0：未婚，1：已婚，2：通用）
	MaritalStatus string `json:"marital_status,omitempty" xml:"marital_status,omitempty"`
	// http://123123.jpg
	DetailImage string `json:"detail_image,omitempty" xml:"detail_image,omitempty"`
	// http://123123.jpg
	ListImage string `json:"list_image,omitempty" xml:"list_image,omitempty"`
	// 市场价（单位分）民营平台商品必需要，并且一个商品组下面的所有套餐市场价必须一样
	MarkPrice string `json:"mark_price,omitempty" xml:"mark_price,omitempty"`
	// http://123123.jpg
	InfoImage string `json:"info_image,omitempty" xml:"info_image,omitempty"`
}
