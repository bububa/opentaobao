package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreReallocateAPIResponse rellocate API返回值
// taobao.omniorder.store.reallocate
//
// 门店发货提供改派接口
type TaobaoOmniorderStoreReallocateAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreReallocateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreReallocateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreReallocateAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreReallocateAPIResponseModel is rellocate 成功返回结果
type TaobaoOmniorderStoreReallocateAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_reallocate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreReallocateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreReallocateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreReallocateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreReallocateAPIResponse)
	},
}

// GetTaobaoOmniorderStoreReallocateAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreReallocateAPIResponse
func GetTaobaoOmniorderStoreReallocateAPIResponse() *TaobaoOmniorderStoreReallocateAPIResponse {
	return poolTaobaoOmniorderStoreReallocateAPIResponse.Get().(*TaobaoOmniorderStoreReallocateAPIResponse)
}

// ReleaseTaobaoOmniorderStoreReallocateAPIResponse 将 TaobaoOmniorderStoreReallocateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreReallocateAPIResponse(v *TaobaoOmniorderStoreReallocateAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreReallocateAPIResponse.Put(v)
}
