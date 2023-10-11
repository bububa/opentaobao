package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeConfirmConsignAPIResponse 换货商家确认收货并发货 API返回值
// tmall.exchange.confirm.consign
//
// 卖家确认收货并发货
type TmallExchangeConfirmConsignAPIResponse struct {
	model.CommonResponse
	TmallExchangeConfirmConsignAPIResponseModel
}

// TmallExchangeConfirmConsignAPIResponseModel is 换货商家确认收货并发货 成功返回结果
type TmallExchangeConfirmConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_confirm_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
