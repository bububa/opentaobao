package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreAccpetedAPIResponse Pos端门店接单接口 API返回值
// taobao.omniorder.store.accpeted
//
// ISV Pos端门店接单，通知星盘
type TaobaoOmniorderStoreAccpetedAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreAccpetedAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreAccpetedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreAccpetedAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreAccpetedAPIResponseModel is Pos端门店接单接口 成功返回结果
type TaobaoOmniorderStoreAccpetedAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_accpeted_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误内容
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreAccpetedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrCode = ""
	m.Message = ""
}

var poolTaobaoOmniorderStoreAccpetedAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreAccpetedAPIResponse)
	},
}

// GetTaobaoOmniorderStoreAccpetedAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreAccpetedAPIResponse
func GetTaobaoOmniorderStoreAccpetedAPIResponse() *TaobaoOmniorderStoreAccpetedAPIResponse {
	return poolTaobaoOmniorderStoreAccpetedAPIResponse.Get().(*TaobaoOmniorderStoreAccpetedAPIResponse)
}

// ReleaseTaobaoOmniorderStoreAccpetedAPIResponse 将 TaobaoOmniorderStoreAccpetedAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreAccpetedAPIResponse(v *TaobaoOmniorderStoreAccpetedAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreAccpetedAPIResponse.Put(v)
}
