package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreConsignedAPIResponse Pos端门店发货 API返回值
// taobao.omniorder.store.consigned
//
// ISV Pos端门店发货，通知星盘
type TaobaoOmniorderStoreConsignedAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreConsignedAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreConsignedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderStoreConsignedAPIResponseModel).Reset()
}

// TaobaoOmniorderStoreConsignedAPIResponseModel is Pos端门店发货 成功返回结果
type TaobaoOmniorderStoreConsignedAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_consigned_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误内容
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *StoreConsignedResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderStoreConsignedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrCode = ""
	m.Message = ""
	m.Data = nil
}

var poolTaobaoOmniorderStoreConsignedAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStoreConsignedAPIResponse)
	},
}

// GetTaobaoOmniorderStoreConsignedAPIResponse 从 sync.Pool 获取 TaobaoOmniorderStoreConsignedAPIResponse
func GetTaobaoOmniorderStoreConsignedAPIResponse() *TaobaoOmniorderStoreConsignedAPIResponse {
	return poolTaobaoOmniorderStoreConsignedAPIResponse.Get().(*TaobaoOmniorderStoreConsignedAPIResponse)
}

// ReleaseTaobaoOmniorderStoreConsignedAPIResponse 将 TaobaoOmniorderStoreConsignedAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderStoreConsignedAPIResponse(v *TaobaoOmniorderStoreConsignedAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderStoreConsignedAPIResponse.Put(v)
}
