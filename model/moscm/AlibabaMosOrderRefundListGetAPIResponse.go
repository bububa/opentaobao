package moscm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosorderrefundlistgetAPIResponse 批量查询交易退货信息 API返回值
// alibaba.mos.order.refund.list.get
//
// 批量查询多个退货单的退货明细
type AlibabamosorderrefundlistgetAPIResponse struct {
	model.CommonResponse
	AlibabamosorderrefundlistgetAPIResponseModel
}

// AlibabamosorderrefundlistgetAPIResponseModel is 批量查询交易退货信息 成功返回结果
type AlibabamosorderrefundlistgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_order_refund_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabamosorderrefundlistgetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
