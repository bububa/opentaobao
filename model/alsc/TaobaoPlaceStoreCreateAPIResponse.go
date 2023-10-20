package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreCreateAPIResponse 商户门店创建接口 API返回值
// taobao.place.store.create
//
// 用于商家创建线下门店
type TaobaoPlaceStoreCreateAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreCreateAPIResponseModel).Reset()
}

// TaobaoPlaceStoreCreateAPIResponseModel is 商户门店创建接口 成功返回结果
type TaobaoPlaceStoreCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StoreId = 0
}

var poolTaobaoPlaceStoreCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreCreateAPIResponse)
	},
}

// GetTaobaoPlaceStoreCreateAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreCreateAPIResponse
func GetTaobaoPlaceStoreCreateAPIResponse() *TaobaoPlaceStoreCreateAPIResponse {
	return poolTaobaoPlaceStoreCreateAPIResponse.Get().(*TaobaoPlaceStoreCreateAPIResponse)
}

// ReleaseTaobaoPlaceStoreCreateAPIResponse 将 TaobaoPlaceStoreCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreCreateAPIResponse(v *TaobaoPlaceStoreCreateAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreCreateAPIResponse.Put(v)
}
