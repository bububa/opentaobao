package examination

// Package 结构体
type Package struct {
	// 体检项列表
	ItemList []Item `json:"item_list,omitempty" xml:"item_list>item,omitempty"`
	// 体检业务：入职体检,父母体检,女性体检,中青年体检,肿瘤筛查,肠胃套餐,健康证,儿童体检,婚前孕前,高端体检,核酸检测,妇科检查,特价体检,证件体检,专科检查,健康优选,减重服务 护理/检测业务：上门换药,陪诊/陪护,小儿护理,日常护理,置管护理,上门核酸,上门体检,传染病检测,慢病检测,肿瘤早筛,甲状腺检测,专项检测,卵巢检测,打针/采血,核酸团检,中医养生,过敏原检测,小儿检测,基因检测,皮肤检测,过敏检测,肠胃检查,甲状腺检查 心理健康业务：心理测评,心理咨询,在线咨询,情感分析,情感顾问,测评,心理医生,咨询,心理辅导,情感咨询,心理电话咨询,心理门店咨询,恋爱情感,焦虑缓解,职业规划,家庭教育,一对一服务,猜你喜欢,推荐,专家课程,线上陪伴营,专家预约,团体治疗,亲子关系,抑郁焦虑 口腔齿科业务：洗牙,儿童口腔,美白,拔牙,烤瓷牙;  多个请用逗号分隔，不建议跨行业/业务使用上述标签；个人到店核酸检测请选择'核酸检测', 上门核酸检测请选择'上门核酸', 上门核酸团检请选择'核酸团检'
	Labels []string `json:"labels,omitempty" xml:"labels>string,omitempty"`
	// 体检业务：三高检测,肿瘤筛查,前列腺检查,宫颈癌筛查,妇科检查,甲状腺检查,心脑血管检查,胸肺部检查,肝胆检查,乳腺检查,颈椎腰椎,肠胃疾病,专项检查,脱发检测,脊椎检测 心理健康业务：1V1专业沟通,专业量表,全面报告,专业心理评估,60分钟深入咨询,个性化定制方案,系统稳定咨询 口腔齿科业务：预防牙周病,清洁牙齿,清除色素,效果更佳,牙齿防蛀,重点保护,促进钙化,全口防蛀,一次搞定,全面检查,赠送X光片,进口树脂,美观耐用,纳米树脂,及时止龋,为换牙护航,微创拔牙,解除隐患,预防感染,进口材料,SPA式体验,不酸不痛,牙齿美白,数字化定制,仿若真牙,认证医生,专业服务; 多个请用逗号分隔，不建议跨行业/业务使用上述标签
	FeatureItem []string `json:"feature_item,omitempty" xml:"feature_item>string,omitempty"`
	// 服务类型 ：ONSITE_SERVICE（到店检测）, DOOR_TO_DOOR_SERVICE（上门检测） ; 多种服务类型用逗号分隔
	ServiceTypes []string `json:"service_types,omitempty" xml:"service_types>string,omitempty"`
	// 关联加项编码
	SkuItemCodes []string `json:"sku_item_codes,omitempty" xml:"sku_item_codes>string,omitempty"`
	// 推荐加项编码
	RecommendSkuItemCodes []string `json:"recommend_sku_item_codes,omitempty" xml:"recommend_sku_item_codes>string,omitempty"`
	// 售卖价（单位分）（数值需要大于100，且大于settlement_price字段的数值），精确到分，民营商品必需要，并且一个商品组下面的所有套餐售卖价必须一样
	ContractPrice string `json:"contract_price,omitempty" xml:"contract_price,omitempty"`
	// 结算价（单位分）（数值需要大于100），精确到分
	SettlementPrice string `json:"settlement_price,omitempty" xml:"settlement_price,omitempty"`
	// 套餐代码，机构保证全局唯一
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 套餐名称（必填！！！）
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 最多500个字，套餐详情
	PackageDetail string `json:"package_detail,omitempty" xml:"package_detail,omitempty"`
	// 性别（0：男，1：女，2：通用）
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
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
	// 报告最大产出时长，服务商承诺，单位小时;0: 无效值，默认为0; 1~4: 4小时出报告, 5~6: 6小时出报告, 7~12: 12小时出报告, 13~24: 24小时出报告
	ReportGenerationMaxDurationBusinessCommitment int64 `json:"report_generation_max_duration_business_commitment,omitempty" xml:"report_generation_max_duration_business_commitment,omitempty"`
	// 是否免费，仅部分场景适用，默认传false
	Free bool `json:"free,omitempty" xml:"free,omitempty"`
	// 是否允许线下支付，仅部分场景适用，默认传false
	AllowOfflinePay bool `json:"allow_offline_pay,omitempty" xml:"allow_offline_pay,omitempty"`
	// 是否只作为加项
	OnlyAsSkuItem bool `json:"only_as_sku_item,omitempty" xml:"only_as_sku_item,omitempty"`
	// 标记加项编码为空，用于清空加项
	SkuItemCodesNoData bool `json:"sku_item_codes_no_data,omitempty" xml:"sku_item_codes_no_data,omitempty"`
}
