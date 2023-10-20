package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatradealiancecreateAPIResponse 推客平台订单回流 API返回值
// alibaba.trade.aliance.create
//
// 推客平台订单回流
type AlibabatradealiancecreateAPIResponse struct {
	model.CommonResponse
	AlibabatradealiancecreateAPIResponseModel
}

// AlibabatradealiancecreateAPIResponseModel is 推客平台订单回流 成功返回结果
type AlibabatradealiancecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_trade_aliance_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单创建结果
	Result *AlibabatradealiancecreateResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
