package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymindustryrecommendgoodsgetAPIResponse 获取交易猫推荐商品 API返回值
// alibaba.jym.industry.recommend.goods.get
//
// 获取交易猫推荐商品
type AlibabajymindustryrecommendgoodsgetAPIResponse struct {
	model.CommonResponse
	AlibabajymindustryrecommendgoodsgetAPIResponseModel
}

// AlibabajymindustryrecommendgoodsgetAPIResponseModel is 获取交易猫推荐商品 成功返回结果
type AlibabajymindustryrecommendgoodsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_industry_recommend_goods_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务编码
	SubCodeType string `json:"sub_code_type,omitempty" xml:"sub_code_type,omitempty"`
	// 业务错误信息
	SubExtraErrMsg string `json:"sub_extra_err_msg,omitempty" xml:"sub_extra_err_msg,omitempty"`
	// 错误码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 系统错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 结果集
	Result *JymRecommendGoodsInfoDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
