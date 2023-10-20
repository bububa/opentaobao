package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtquerystationAPIResponse 速店通查询站点信息 API返回值
// taobao.omniorder.store.sdtquerystation
//
// 速店通查询站点信息
type TaobaoOmniorderStoreSdtquerystationAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreSdtquerystationAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSdtquerystationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreSdtquerystationAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreSdtquerystationAPIResponseModel is 速店通查询站点信息 成功返回结果
type TaobaoOmniorderStoreSdtquerystationAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_sdtquerystation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreSdtquerystationResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreSdtquerystationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniorderStoreSdtquerystationAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreSdtquerystationAPIResponse)
	},
}

// GetTaobaoOmniorderStoreSdtquerystationAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreSdtquerystationAPIResponse
func GetTaobaoOmniorderStoreSdtquerystationAPIResponse() *TaobaoOmniorderStoreSdtquerystationAPIResponse {
	return poolTaobaoOmniorderStoreSdtquerystationAPIResponse.Get().(*TaobaoOmniorderStoreSdtquerystationAPIResponse)
}

// ReleaseTaobaoOmniorderStoreSdtquerystationAPIResponse 将 TaobaoOmniorderStoreSdtquerystationAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreSdtquerystationAPIResponse(v *TaobaoOmniorderStoreSdtquerystationAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreSdtquerystationAPIResponse.Put(v)
}
