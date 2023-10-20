package xhotelofficial

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfficialSettleCancelAPIResponse 官网信用住取消结账 API返回值
// taobao.xhotel.order.official.settle.cancel
//
// 用于官网信用住取消结账
type TaobaoXhotelOrderOfficialSettleCancelAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderOfficialSettleCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderOfficialSettleCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderOfficialSettleCancelAPIResponseModel).Reset()
}

// TaobaoXhotelOrderOfficialSettleCancelAPIResponseModel is 官网信用住取消结账 成功返回结果
type TaobaoXhotelOrderOfficialSettleCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_official_settle_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderOfficialSettleCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderOfficialSettleCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderOfficialSettleCancelAPIResponse)
	},
}

// GetTaobaoXhotelOrderOfficialSettleCancelAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderOfficialSettleCancelAPIResponse
func GetTaobaoXhotelOrderOfficialSettleCancelAPIResponse() *TaobaoXhotelOrderOfficialSettleCancelAPIResponse {
	return poolTaobaoXhotelOrderOfficialSettleCancelAPIResponse.Get().(*TaobaoXhotelOrderOfficialSettleCancelAPIResponse)
}

// ReleaseTaobaoXhotelOrderOfficialSettleCancelAPIResponse 将 TaobaoXhotelOrderOfficialSettleCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderOfficialSettleCancelAPIResponse(v *TaobaoXhotelOrderOfficialSettleCancelAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderOfficialSettleCancelAPIResponse.Put(v)
}
