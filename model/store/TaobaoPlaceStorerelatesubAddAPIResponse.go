package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorerelatesubAddAPIResponse 门店和子门店关系新增 API返回值
// taobao.place.storerelatesub.add
//
// 门店和子门店关系新增
type TaobaoPlaceStorerelatesubAddAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStorerelatesubAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStorerelatesubAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStorerelatesubAddAPIResponseModel).Reset()
}

// TaobaoPlaceStorerelatesubAddAPIResponseModel is 门店和子门店关系新增 成功返回结果
type TaobaoPlaceStorerelatesubAddAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storerelatesub_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStorerelatesubAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStorerelatesubAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStorerelatesubAddAPIResponse)
	},
}

// GetTaobaoPlaceStorerelatesubAddAPIResponse 从 sync.Pool 获取 TaobaoPlaceStorerelatesubAddAPIResponse
func GetTaobaoPlaceStorerelatesubAddAPIResponse() *TaobaoPlaceStorerelatesubAddAPIResponse {
	return poolTaobaoPlaceStorerelatesubAddAPIResponse.Get().(*TaobaoPlaceStorerelatesubAddAPIResponse)
}

// ReleaseTaobaoPlaceStorerelatesubAddAPIResponse 将 TaobaoPlaceStorerelatesubAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStorerelatesubAddAPIResponse(v *TaobaoPlaceStorerelatesubAddAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStorerelatesubAddAPIResponse.Put(v)
}
