package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlifanyiMarketAuthenticateAPIResponse 第三方授权 API返回值
// alibaba.alifanyi.market.authenticate
//
// 第三方授权，获取授权码
type AlibabaAlifanyiMarketAuthenticateAPIResponse struct {
	model.CommonResponse
	AlibabaAlifanyiMarketAuthenticateAPIResponseModel
}

// AlibabaAlifanyiMarketAuthenticateAPIResponseModel is 第三方授权 成功返回结果
type AlibabaAlifanyiMarketAuthenticateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alifanyi_market_authenticate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权码
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	BizErrorCode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 接口是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
