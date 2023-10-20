package xhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfflineSettleCancelAPIResponse 线下信用住取消结账专用接口 API返回值
// taobao.xhotel.order.offline.settle.cancel
//
// 线下信用住取消结账专用接口
type TaobaoXhotelOrderOfflineSettleCancelAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderOfflineSettleCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderOfflineSettleCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderOfflineSettleCancelAPIResponseModel).Reset()
}

// TaobaoXhotelOrderOfflineSettleCancelAPIResponseModel is 线下信用住取消结账专用接口 成功返回结果
type TaobaoXhotelOrderOfflineSettleCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_offline_settle_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderOfflineSettleCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderOfflineSettleCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderOfflineSettleCancelAPIResponse)
	},
}

// GetTaobaoXhotelOrderOfflineSettleCancelAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderOfflineSettleCancelAPIResponse
func GetTaobaoXhotelOrderOfflineSettleCancelAPIResponse() *TaobaoXhotelOrderOfflineSettleCancelAPIResponse {
	return poolTaobaoXhotelOrderOfflineSettleCancelAPIResponse.Get().(*TaobaoXhotelOrderOfflineSettleCancelAPIResponse)
}

// ReleaseTaobaoXhotelOrderOfflineSettleCancelAPIResponse 将 TaobaoXhotelOrderOfflineSettleCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderOfflineSettleCancelAPIResponse(v *TaobaoXhotelOrderOfflineSettleCancelAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderOfflineSettleCancelAPIResponse.Put(v)
}
