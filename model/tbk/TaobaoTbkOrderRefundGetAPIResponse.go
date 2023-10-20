package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkOrderRefundGetAPIResponse 淘宝客-推广者-全量维权退款订单查询 API返回值
// taobao.tbk.order.refund.get
//
// 淘宝客账户下全量维权退款订单查询
type TaobaoTbkOrderRefundGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkOrderRefundGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkOrderRefundGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkOrderRefundGetAPIResponseModel).Reset()
}

// TaobaoTbkOrderRefundGetAPIResponseModel is 淘宝客-推广者-全量维权退款订单查询 成功返回结果
type TaobaoTbkOrderRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_order_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务错误信息
	BizErrorDesc string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
	// 返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回数据
	Data *OrderPage `json:"data,omitempty" xml:"data,omitempty"`
	// 接口返回值信息，跟rpc架构保持一致
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 业务错误码 101, 102,103
	BizErrorCode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkOrderRefundGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorDesc = ""
	m.ResultMsg = ""
	m.Data = nil
	m.ResultCode = 0
	m.BizErrorCode = 0
}

var poolTaobaoTbkOrderRefundGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkOrderRefundGetAPIResponse)
	},
}

// GetTaobaoTbkOrderRefundGetAPIResponse 从 sync.Pool 获取 TaobaoTbkOrderRefundGetAPIResponse
func GetTaobaoTbkOrderRefundGetAPIResponse() *TaobaoTbkOrderRefundGetAPIResponse {
	return poolTaobaoTbkOrderRefundGetAPIResponse.Get().(*TaobaoTbkOrderRefundGetAPIResponse)
}

// ReleaseTaobaoTbkOrderRefundGetAPIResponse 将 TaobaoTbkOrderRefundGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkOrderRefundGetAPIResponse(v *TaobaoTbkOrderRefundGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkOrderRefundGetAPIResponse.Put(v)
}
