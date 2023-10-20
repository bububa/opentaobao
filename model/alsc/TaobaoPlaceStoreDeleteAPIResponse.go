package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreDeleteAPIResponse 线下门店删除 API返回值
// taobao.place.store.delete
//
// 用于商家删除线下门店
type TaobaoPlaceStoreDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreDeleteAPIResponseModel).Reset()
}

// TaobaoPlaceStoreDeleteAPIResponseModel is 线下门店删除 成功返回结果
type TaobaoPlaceStoreDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店删除结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoPlaceStoreDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreDeleteAPIResponse)
	},
}

// GetTaobaoPlaceStoreDeleteAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreDeleteAPIResponse
func GetTaobaoPlaceStoreDeleteAPIResponse() *TaobaoPlaceStoreDeleteAPIResponse {
	return poolTaobaoPlaceStoreDeleteAPIResponse.Get().(*TaobaoPlaceStoreDeleteAPIResponse)
}

// ReleaseTaobaoPlaceStoreDeleteAPIResponse 将 TaobaoPlaceStoreDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreDeleteAPIResponse(v *TaobaoPlaceStoreDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreDeleteAPIResponse.Put(v)
}
