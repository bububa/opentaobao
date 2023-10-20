package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoreceiptordercheckAPIResponse 查看工单查询订单是否已付款 API返回值
// tmall.aliauto.receipt.order.check
//
// 查看工单查询订单是否已付款
type TmallaliautoreceiptordercheckAPIResponse struct {
	model.CommonResponse
	TmallaliautoreceiptordercheckAPIResponseModel
}

// TmallaliautoreceiptordercheckAPIResponseModel is 查看工单查询订单是否已付款 成功返回结果
type TmallaliautoreceiptordercheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_receipt_order_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *CheckReceiptOrderIsPaid4isvDto `json:"data,omitempty" xml:"data,omitempty"`
}
