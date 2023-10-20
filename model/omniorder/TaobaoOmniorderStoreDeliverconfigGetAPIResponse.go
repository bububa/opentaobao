package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreDeliverconfigGetAPIResponse 查询门店发货配置内容 API返回值
// taobao.omniorder.store.deliverconfig.get
//
// 查询门店发货配置内容
type TaobaoOmniorderStoreDeliverconfigGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreDeliverconfigGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel is 查询门店发货配置内容 成功返回结果
type TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_deliverconfig_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreDeliverconfigGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreDeliverconfigGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreDeliverconfigGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreDeliverconfigGetAPIResponse)
	},
}

// GetTaobaoOmniorderStoreDeliverconfigGetAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreDeliverconfigGetAPIResponse
func GetTaobaoOmniorderStoreDeliverconfigGetAPIResponse() *TaobaoOmniorderStoreDeliverconfigGetAPIResponse {
	return poolTaobaoOmniorderStoreDeliverconfigGetAPIResponse.Get().(*TaobaoOmniorderStoreDeliverconfigGetAPIResponse)
}

// ReleaseTaobaoOmniorderStoreDeliverconfigGetAPIResponse 将 TaobaoOmniorderStoreDeliverconfigGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreDeliverconfigGetAPIResponse(v *TaobaoOmniorderStoreDeliverconfigGetAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreDeliverconfigGetAPIResponse.Put(v)
}
