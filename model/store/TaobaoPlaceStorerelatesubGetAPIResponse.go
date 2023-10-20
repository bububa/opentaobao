package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorerelatesubGetAPIResponse 门店和子门店关系查找 API返回值
// taobao.place.storerelatesub.get
//
// 门店和子门店关系查找
type TaobaoPlaceStorerelatesubGetAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStorerelatesubGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStorerelatesubGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStorerelatesubGetAPIResponseModel).Reset()
}

// TaobaoPlaceStorerelatesubGetAPIResponseModel is 门店和子门店关系查找 成功返回结果
type TaobaoPlaceStorerelatesubGetAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storerelatesub_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStorerelatesubGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStorerelatesubGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStorerelatesubGetAPIResponse)
	},
}

// GetTaobaoPlaceStorerelatesubGetAPIResponse 从 sync.Pool 获取 TaobaoPlaceStorerelatesubGetAPIResponse
func GetTaobaoPlaceStorerelatesubGetAPIResponse() *TaobaoPlaceStorerelatesubGetAPIResponse {
	return poolTaobaoPlaceStorerelatesubGetAPIResponse.Get().(*TaobaoPlaceStorerelatesubGetAPIResponse)
}

// ReleaseTaobaoPlaceStorerelatesubGetAPIResponse 将 TaobaoPlaceStorerelatesubGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStorerelatesubGetAPIResponse(v *TaobaoPlaceStorerelatesubGetAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStorerelatesubGetAPIResponse.Put(v)
}
