package xhoteloffline

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfflineSettlePutAPIResponse 线下信用住结账专用接口 API返回值
// taobao.xhotel.order.offline.settle.put
//
// 线下信用住结账专用接口
type TaobaoXhotelOrderOfflineSettlePutAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderOfflineSettlePutAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderOfflineSettlePutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderOfflineSettlePutAPIResponseModel).Reset()
}

// TaobaoXhotelOrderOfflineSettlePutAPIResponseModel is 线下信用住结账专用接口 成功返回结果
type TaobaoXhotelOrderOfflineSettlePutAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_offline_settle_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderOfflineSettlePutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderOfflineSettlePutAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderOfflineSettlePutAPIResponse)
	},
}

// GetTaobaoXhotelOrderOfflineSettlePutAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderOfflineSettlePutAPIResponse
func GetTaobaoXhotelOrderOfflineSettlePutAPIResponse() *TaobaoXhotelOrderOfflineSettlePutAPIResponse {
	return poolTaobaoXhotelOrderOfflineSettlePutAPIResponse.Get().(*TaobaoXhotelOrderOfflineSettlePutAPIResponse)
}

// ReleaseTaobaoXhotelOrderOfflineSettlePutAPIResponse 将 TaobaoXhotelOrderOfflineSettlePutAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderOfflineSettlePutAPIResponse(v *TaobaoXhotelOrderOfflineSettlePutAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderOfflineSettlePutAPIResponse.Put(v)
}
