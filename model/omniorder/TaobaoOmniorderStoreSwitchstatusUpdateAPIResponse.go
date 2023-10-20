package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse switchstatus.update API返回值
// taobao.omniorder.store.switchstatus.update
//
// 变更门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreSwitchstatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreSwitchstatusUpdateAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreSwitchstatusUpdateAPIResponseModel is switchstatus.update 成功返回结果
type TaobaoOmniorderStoreSwitchstatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_switchstatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreSwitchstatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSwitchstatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreSwitchstatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse)
	},
}

// GetTaobaoOmniorderStoreSwitchstatusUpdateAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse
func GetTaobaoOmniorderStoreSwitchstatusUpdateAPIResponse() *TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse {
	return poolTaobaoOmniorderStoreSwitchstatusUpdateAPIResponse.Get().(*TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse)
}

// ReleaseTaobaoOmniorderStoreSwitchstatusUpdateAPIResponse 将 TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreSwitchstatusUpdateAPIResponse(v *TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreSwitchstatusUpdateAPIResponse.Put(v)
}
