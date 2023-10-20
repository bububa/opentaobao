package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoOrderConfirmPaidAPIResponse 确认收款 API返回值
// taobao.fenxiao.order.confirm.paid
//
// 供应商确认收款（非支付宝交易）。
type TaobaoFenxiaoOrderConfirmPaidAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoOrderConfirmPaidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoOrderConfirmPaidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoOrderConfirmPaidAPIResponseModel).Reset()
}

// TaobaoFenxiaoOrderConfirmPaidAPIResponseModel is 确认收款 成功返回结果
type TaobaoFenxiaoOrderConfirmPaidAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_order_confirm_paid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 确认结果成功与否
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoOrderConfirmPaidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoOrderConfirmPaidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoOrderConfirmPaidAPIResponse)
	},
}

// GetTaobaoFenxiaoOrderConfirmPaidAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoOrderConfirmPaidAPIResponse
func GetTaobaoFenxiaoOrderConfirmPaidAPIResponse() *TaobaoFenxiaoOrderConfirmPaidAPIResponse {
	return poolTaobaoFenxiaoOrderConfirmPaidAPIResponse.Get().(*TaobaoFenxiaoOrderConfirmPaidAPIResponse)
}

// ReleaseTaobaoFenxiaoOrderConfirmPaidAPIResponse 将 TaobaoFenxiaoOrderConfirmPaidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoOrderConfirmPaidAPIResponse(v *TaobaoFenxiaoOrderConfirmPaidAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoOrderConfirmPaidAPIResponse.Put(v)
}
