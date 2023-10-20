package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreQueryAPIResponse 门店信息查询接口 API返回值
// taobao.place.store.query
//
// 根据用户授权信息，获取用户的门店公开信息
type TaobaoPlaceStoreQueryAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreQueryAPIResponseModel).Reset()
}

// TaobaoPlaceStoreQueryAPIResponseModel is 门店信息查询接口 成功返回结果
type TaobaoPlaceStoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStoreQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreQueryAPIResponse)
	},
}

// GetTaobaoPlaceStoreQueryAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreQueryAPIResponse
func GetTaobaoPlaceStoreQueryAPIResponse() *TaobaoPlaceStoreQueryAPIResponse {
	return poolTaobaoPlaceStoreQueryAPIResponse.Get().(*TaobaoPlaceStoreQueryAPIResponse)
}

// ReleaseTaobaoPlaceStoreQueryAPIResponse 将 TaobaoPlaceStoreQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreQueryAPIResponse(v *TaobaoPlaceStoreQueryAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreQueryAPIResponse.Put(v)
}
