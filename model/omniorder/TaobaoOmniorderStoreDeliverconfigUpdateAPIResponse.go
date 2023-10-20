package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse 修改门店发货配置内容 API返回值
// taobao.omniorder.store.deliverconfig.update
//
// 修改门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreDeliverconfigUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreDeliverconfigUpdateAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreDeliverconfigUpdateAPIResponseModel is 修改门店发货配置内容 成功返回结果
type TaobaoOmniorderStoreDeliverconfigUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_deliverconfig_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreDeliverconfigUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreDeliverconfigUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreDeliverconfigUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse)
	},
}

// GetTaobaoOmniorderStoreDeliverconfigUpdateAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse
func GetTaobaoOmniorderStoreDeliverconfigUpdateAPIResponse() *TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse {
	return poolTaobaoOmniorderStoreDeliverconfigUpdateAPIResponse.Get().(*TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse)
}

// ReleaseTaobaoOmniorderStoreDeliverconfigUpdateAPIResponse 将 TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreDeliverconfigUpdateAPIResponse(v *TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreDeliverconfigUpdateAPIResponse.Put(v)
}
