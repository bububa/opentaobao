package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreprocessCreateAPIResponse 仓内加工单创建接口 API返回值
// taobao.qimen.storeprocess.create
//
// ERP调用奇门的接口,创建仓内加工单
type TaobaoQimenStoreprocessCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStoreprocessCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStoreprocessCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStoreprocessCreateAPIResponseModel).Reset()
}

// TaobaoQimenStoreprocessCreateAPIResponseModel is 仓内加工单创建接口 成功返回结果
type TaobaoQimenStoreprocessCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_storeprocess_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *StoreProcessCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStoreprocessCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenStoreprocessCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStoreprocessCreateAPIResponse)
	},
}

// GetTaobaoQimenStoreprocessCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenStoreprocessCreateAPIResponse
func GetTaobaoQimenStoreprocessCreateAPIResponse() *TaobaoQimenStoreprocessCreateAPIResponse {
	return poolTaobaoQimenStoreprocessCreateAPIResponse.Get().(*TaobaoQimenStoreprocessCreateAPIResponse)
}

// ReleaseTaobaoQimenStoreprocessCreateAPIResponse 将 TaobaoQimenStoreprocessCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStoreprocessCreateAPIResponse(v *TaobaoQimenStoreprocessCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenStoreprocessCreateAPIResponse.Put(v)
}
