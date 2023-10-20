package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryTradeMaxPriceGetAPIResponse 获取交易猫单个游戏渠道帐号交易成功最高价 API返回值
// alibaba.jym.industry.trade.max.price.get
//
// 获取交易猫单个游戏渠道帐号交易成功最高价
type AlibabaJymIndustryTradeMaxPriceGetAPIResponse struct {
	model.CommonResponse
	AlibabaJymIndustryTradeMaxPriceGetAPIResponseModel
}

// AlibabaJymIndustryTradeMaxPriceGetAPIResponseModel is 获取交易猫单个游戏渠道帐号交易成功最高价 成功返回结果
type AlibabaJymIndustryTradeMaxPriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_industry_trade_max_price_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子错误码
	SubCodeType string `json:"sub_code_type,omitempty" xml:"sub_code_type,omitempty"`
	// 子错误码描述
	SubExtraErrMsg string `json:"sub_extra_err_msg,omitempty" xml:"sub_extra_err_msg,omitempty"`
	// 错误码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 最高价订单信息
	Result *JymMaxPriceOrderInfoDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
