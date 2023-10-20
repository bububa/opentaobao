package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsReturnBillGetAPIResponse 获取销退收货信息 API返回值
// taobao.wlb.wms.return.bill.get
//
// 通过订单号获取单个销退单收货信息。
type TaobaoWlbWmsReturnBillGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsReturnBillGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsReturnBillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsReturnBillGetAPIResponseModel).Reset()
}

// TaobaoWlbWmsReturnBillGetAPIResponseModel is 获取销退收货信息 成功返回结果
type TaobaoWlbWmsReturnBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_return_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回退订单信息
	ReturnOrderInfo *CainiaoReturnBillReturnorderinfo `json:"return_order_info,omitempty" xml:"return_order_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsReturnBillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnOrderInfo = nil
}

var poolTaobaoWlbWmsReturnBillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsReturnBillGetAPIResponse)
	},
}

// GetTaobaoWlbWmsReturnBillGetAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsReturnBillGetAPIResponse
func GetTaobaoWlbWmsReturnBillGetAPIResponse() *TaobaoWlbWmsReturnBillGetAPIResponse {
	return poolTaobaoWlbWmsReturnBillGetAPIResponse.Get().(*TaobaoWlbWmsReturnBillGetAPIResponse)
}

// ReleaseTaobaoWlbWmsReturnBillGetAPIResponse 将 TaobaoWlbWmsReturnBillGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsReturnBillGetAPIResponse(v *TaobaoWlbWmsReturnBillGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsReturnBillGetAPIResponse.Put(v)
}
