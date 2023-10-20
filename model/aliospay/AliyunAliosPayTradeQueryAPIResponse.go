package aliospay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayTradeQueryAPIResponse 查询支付结果接口 API返回值
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
type AliyunAliosPayTradeQueryAPIResponse struct {
	model.CommonResponse
	AliyunAliosPayTradeQueryAPIResponseModel
}

// AliyunAliosPayTradeQueryAPIResponseModel is 查询支付结果接口 成功返回结果
type AliyunAliosPayTradeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_trade_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOspayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}
