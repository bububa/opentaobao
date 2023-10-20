package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtstatusAPIResponse 菜鸟裹裹运单状态查询 API返回值
// taobao.omniorder.store.sdtstatus
//
// 提供给商家查询运力单的状态。
type TaobaoOmniorderStoreSdtstatusAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreSdtstatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSdtstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreSdtstatusAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreSdtstatusAPIResponseModel is 菜鸟裹裹运单状态查询 成功返回结果
type TaobaoOmniorderStoreSdtstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_sdtstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreSdtstatusResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSdtstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreSdtstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSdtstatusAPIResponse)
	},
}

// GetTaobaoOmniorderStoreSdtstatusAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreSdtstatusAPIResponse
func GetTaobaoOmniorderStoreSdtstatusAPIResponse() *TaobaoOmniorderStoreSdtstatusAPIResponse {
	return poolTaobaoOmniorderStoreSdtstatusAPIResponse.Get().(*TaobaoOmniorderStoreSdtstatusAPIResponse)
}

// ReleaseTaobaoOmniorderStoreSdtstatusAPIResponse 将 TaobaoOmniorderStoreSdtstatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreSdtstatusAPIResponse(v *TaobaoOmniorderStoreSdtstatusAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreSdtstatusAPIResponse.Put(v)
}
