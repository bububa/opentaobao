package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyOrderReceiptQueryAPIResponse 订单小票查询 API返回值
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
type AlibabaTclsAelophyOrderReceiptQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyOrderReceiptQueryAPIResponseModel
}

// AlibabaTclsAelophyOrderReceiptQueryAPIResponseModel is 订单小票查询 成功返回结果
type AlibabaTclsAelophyOrderReceiptQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_order_receipt_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaTclsAelophyOrderReceiptQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
