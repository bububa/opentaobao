package bus

// InsuranceProductVO 结构体
type InsuranceProductVO struct {
	// 保险产品码
	InsProductCode string `json:"ins_product_code,omitempty" xml:"ins_product_code,omitempty"`
	// 保险名称
	InsName string `json:"ins_name,omitempty" xml:"ins_name,omitempty"`
	// 保险标题
	InsTitle string `json:"ins_title,omitempty" xml:"ins_title,omitempty"`
	// 特别约定
	SpecialTermsAndConditions *InsurancePropertyVO `json:"special_terms_and_conditions,omitempty" xml:"special_terms_and_conditions,omitempty"`
	// 产品特色图片
	ProductFeatureImage *InsurancePropertyVO `json:"product_feature_image,omitempty" xml:"product_feature_image,omitempty"`
	// 查看详情
	CheckDtails *InsurancePropertyVO `json:"check_dtails,omitempty" xml:"check_dtails,omitempty"`
	// 理赔流程
	ClaimFlow *InsurancePropertyVO `json:"claim_flow,omitempty" xml:"claim_flow,omitempty"`
	// 保险条款
	ProductTerm *InsurancePropertyVO `json:"product_term,omitempty" xml:"product_term,omitempty"`
	// 产品标签
	ProdTag *InsurancePropertyVO `json:"prod_tag,omitempty" xml:"prod_tag,omitempty"`
	// 产品特色
	ProductFeature *InsurancePropertyVO `json:"product_feature,omitempty" xml:"product_feature,omitempty"`
	// 理赔服务电话
	ClaimServicePhone *InsurancePropertyVO `json:"claim_service_phone,omitempty" xml:"claim_service_phone,omitempty"`
	// 产品图片
	ProductImage *InsurancePropertyVO `json:"product_image,omitempty" xml:"product_image,omitempty"`
	// 投保须知
	InsMustKnow *InsurancePropertyVO `json:"ins_must_know,omitempty" xml:"ins_must_know,omitempty"`
	// 客服电话
	CustomerServicePhone *InsurancePropertyVO `json:"customer_service_phone,omitempty" xml:"customer_service_phone,omitempty"`
	// 保险售卖价格(单位：分)
	InsPrice int64 `json:"ins_price,omitempty" xml:"ins_price,omitempty"`
}
