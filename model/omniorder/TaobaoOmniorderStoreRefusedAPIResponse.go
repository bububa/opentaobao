package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreRefusedAPIResponse Pos端门店拒单 API返回值
// taobao.omniorder.store.refused
//
// ISV Pos端门店拒单，通知星盘
type TaobaoOmniorderStoreRefusedAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreRefusedAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreRefusedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreRefusedAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreRefusedAPIResponseModel is Pos端门店拒单 成功返回结果
type TaobaoOmniorderStoreRefusedAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_refused_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 正常为0,其他表示异常
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreRefusedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrCode = ""
	m.Message = ""
}

var poolTaobaoOmniorderStoreRefusedAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreRefusedAPIResponse)
	},
}

// GetTaobaoOmniorderStoreRefusedAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreRefusedAPIResponse
func GetTaobaoOmniorderStoreRefusedAPIResponse() *TaobaoOmniorderStoreRefusedAPIResponse {
	return poolTaobaoOmniorderStoreRefusedAPIResponse.Get().(*TaobaoOmniorderStoreRefusedAPIResponse)
}

// ReleaseTaobaoOmniorderStoreRefusedAPIResponse 将 TaobaoOmniorderStoreRefusedAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreRefusedAPIResponse(v *TaobaoOmniorderStoreRefusedAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreRefusedAPIResponse.Put(v)
}
