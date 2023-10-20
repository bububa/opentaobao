package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundDetailGetAPIResponse 退款详情页渲染 API返回值
// taobao.refund.detail.get
//
// 退款详情页渲染
type TaobaoRefundDetailGetAPIResponse struct {
	model.CommonResponse
	TaobaoRefundDetailGetAPIResponseModel
}

// TaobaoRefundDetailGetAPIResponseModel is 退款详情页渲染 成功返回结果
type TaobaoRefundDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款详情
	Detail *RefundDetail `json:"detail,omitempty" xml:"detail,omitempty"`
}
