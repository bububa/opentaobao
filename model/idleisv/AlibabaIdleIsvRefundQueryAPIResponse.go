package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvrefundqueryAPIResponse 闲鱼已验货交易订单退款信息查询 API返回值
// alibaba.idle.isv.refund.query
//
// 闲鱼服务商交易订单退款信息查询-单个退款查询
type AlibabaidleisvrefundqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvrefundqueryAPIResponseModel
}

// AlibabaidleisvrefundqueryAPIResponseModel is 闲鱼已验货交易订单退款信息查询 成功返回结果
type AlibabaidleisvrefundqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_refund_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaidleisvrefundqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
