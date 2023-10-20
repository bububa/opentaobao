package bus

import (
	"sync"
)

// InsuranceProductVo 结构体
type InsuranceProductVo struct {
	// 保险产品码
	InsProductCode string `json:"ins_product_code,omitempty" xml:"ins_product_code,omitempty"`
	// 保险名称
	InsName string `json:"ins_name,omitempty" xml:"ins_name,omitempty"`
	// 保险标题
	InsTitle string `json:"ins_title,omitempty" xml:"ins_title,omitempty"`
	// 特别约定
	SpecialTermsAndConditions *InsurancePropertyVo `json:"special_terms_and_conditions,omitempty" xml:"special_terms_and_conditions,omitempty"`
	// 产品特色图片
	ProductFeatureImage *InsurancePropertyVo `json:"product_feature_image,omitempty" xml:"product_feature_image,omitempty"`
	// 查看详情
	CheckDtails *InsurancePropertyVo `json:"check_dtails,omitempty" xml:"check_dtails,omitempty"`
	// 理赔流程
	ClaimFlow *InsurancePropertyVo `json:"claim_flow,omitempty" xml:"claim_flow,omitempty"`
	// 保险条款
	ProductTerm *InsurancePropertyVo `json:"product_term,omitempty" xml:"product_term,omitempty"`
	// 产品标签
	ProdTag *InsurancePropertyVo `json:"prod_tag,omitempty" xml:"prod_tag,omitempty"`
	// 产品特色
	ProductFeature *InsurancePropertyVo `json:"product_feature,omitempty" xml:"product_feature,omitempty"`
	// 理赔服务电话
	ClaimServicePhone *InsurancePropertyVo `json:"claim_service_phone,omitempty" xml:"claim_service_phone,omitempty"`
	// 产品图片
	ProductImage *InsurancePropertyVo `json:"product_image,omitempty" xml:"product_image,omitempty"`
	// 投保须知
	InsMustKnow *InsurancePropertyVo `json:"ins_must_know,omitempty" xml:"ins_must_know,omitempty"`
	// 客服电话
	CustomerServicePhone *InsurancePropertyVo `json:"customer_service_phone,omitempty" xml:"customer_service_phone,omitempty"`
	// 保险售卖价格(单位：分)
	InsPrice int64 `json:"ins_price,omitempty" xml:"ins_price,omitempty"`
}

var poolInsuranceProductVo = sync.Pool{
	New: func() any {
		return new(InsuranceProductVo)
	},
}

// GetInsuranceProductVo() 从对象池中获取InsuranceProductVo
func GetInsuranceProductVo() *InsuranceProductVo {
	return poolInsuranceProductVo.Get().(*InsuranceProductVo)
}

// ReleaseInsuranceProductVo 释放InsuranceProductVo
func ReleaseInsuranceProductVo(v *InsuranceProductVo) {
	v.InsProductCode = ""
	v.InsName = ""
	v.InsTitle = ""
	v.SpecialTermsAndConditions = nil
	v.ProductFeatureImage = nil
	v.CheckDtails = nil
	v.ClaimFlow = nil
	v.ProductTerm = nil
	v.ProdTag = nil
	v.ProductFeature = nil
	v.ClaimServicePhone = nil
	v.ProductImage = nil
	v.InsMustKnow = nil
	v.CustomerServicePhone = nil
	v.InsPrice = 0
	poolInsuranceProductVo.Put(v)
}
