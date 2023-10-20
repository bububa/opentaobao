package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreCollectconfigGetAPIResponse 查询门店自提配置内容 API返回值
// taobao.omniorder.store.collectconfig.get
//
// 查询门店自提配置内容
type TaobaoOmniorderStoreCollectconfigGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreCollectconfigGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreCollectconfigGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreCollectconfigGetAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreCollectconfigGetAPIResponseModel is 查询门店自提配置内容 成功返回结果
type TaobaoOmniorderStoreCollectconfigGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_collectconfig_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreCollectconfigGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreCollectconfigGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreCollectconfigGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreCollectconfigGetAPIResponse)
	},
}

// GetTaobaoOmniorderStoreCollectconfigGetAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreCollectconfigGetAPIResponse
func GetTaobaoOmniorderStoreCollectconfigGetAPIResponse() *TaobaoOmniorderStoreCollectconfigGetAPIResponse {
	return poolTaobaoOmniorderStoreCollectconfigGetAPIResponse.Get().(*TaobaoOmniorderStoreCollectconfigGetAPIResponse)
}

// ReleaseTaobaoOmniorderStoreCollectconfigGetAPIResponse 将 TaobaoOmniorderStoreCollectconfigGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreCollectconfigGetAPIResponse(v *TaobaoOmniorderStoreCollectconfigGetAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreCollectconfigGetAPIResponse.Put(v)
}
