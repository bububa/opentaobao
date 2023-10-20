package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSwitchstatusGetAPIResponse switchstatus.get API返回值
// taobao.omniorder.store.switchstatus.get
//
// 查询门店发货、门店自提状态
type TaobaoOmniorderStoreSwitchstatusGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreSwitchstatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSwitchstatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreSwitchstatusGetAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreSwitchstatusGetAPIResponseModel is switchstatus.get 成功返回结果
type TaobaoOmniorderStoreSwitchstatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_switchstatus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreSwitchstatusGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSwitchstatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreSwitchstatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSwitchstatusGetAPIResponse)
	},
}

// GetTaobaoOmniorderStoreSwitchstatusGetAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreSwitchstatusGetAPIResponse
func GetTaobaoOmniorderStoreSwitchstatusGetAPIResponse() *TaobaoOmniorderStoreSwitchstatusGetAPIResponse {
	return poolTaobaoOmniorderStoreSwitchstatusGetAPIResponse.Get().(*TaobaoOmniorderStoreSwitchstatusGetAPIResponse)
}

// ReleaseTaobaoOmniorderStoreSwitchstatusGetAPIResponse 将 TaobaoOmniorderStoreSwitchstatusGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreSwitchstatusGetAPIResponse(v *TaobaoOmniorderStoreSwitchstatusGetAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreSwitchstatusGetAPIResponse.Put(v)
}
