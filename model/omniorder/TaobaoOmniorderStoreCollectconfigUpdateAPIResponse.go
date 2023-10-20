package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreCollectconfigUpdateAPIResponse 门店自提配置修改 API返回值
// taobao.omniorder.store.collectconfig.update
//
// 修改门店自提配置内容
type TaobaoOmniorderStoreCollectconfigUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreCollectconfigUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreCollectconfigUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreCollectconfigUpdateAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreCollectconfigUpdateAPIResponseModel is 门店自提配置修改 成功返回结果
type TaobaoOmniorderStoreCollectconfigUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_collectconfig_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreCollectconfigUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreCollectconfigUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreCollectconfigUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreCollectconfigUpdateAPIResponse)
	},
}

// GetTaobaoOmniorderStoreCollectconfigUpdateAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreCollectconfigUpdateAPIResponse
func GetTaobaoOmniorderStoreCollectconfigUpdateAPIResponse() *TaobaoOmniorderStoreCollectconfigUpdateAPIResponse {
	return poolTaobaoOmniorderStoreCollectconfigUpdateAPIResponse.Get().(*TaobaoOmniorderStoreCollectconfigUpdateAPIResponse)
}

// ReleaseTaobaoOmniorderStoreCollectconfigUpdateAPIResponse 将 TaobaoOmniorderStoreCollectconfigUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreCollectconfigUpdateAPIResponse(v *TaobaoOmniorderStoreCollectconfigUpdateAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreCollectconfigUpdateAPIResponse.Put(v)
}
