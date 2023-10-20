package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorerelatesubDeleteAPIResponse 门店和子门店关系删除 API返回值
// taobao.place.storerelatesub.delete
//
// 门店和子门店关系删除
type TaobaoPlaceStorerelatesubDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStorerelatesubDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStorerelatesubDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStorerelatesubDeleteAPIResponseModel).Reset()
}

// TaobaoPlaceStorerelatesubDeleteAPIResponseModel is 门店和子门店关系删除 成功返回结果
type TaobaoPlaceStorerelatesubDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storerelatesub_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStorerelatesubDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStorerelatesubDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStorerelatesubDeleteAPIResponse)
	},
}

// GetTaobaoPlaceStorerelatesubDeleteAPIResponse 从 sync.Pool 获取 TaobaoPlaceStorerelatesubDeleteAPIResponse
func GetTaobaoPlaceStorerelatesubDeleteAPIResponse() *TaobaoPlaceStorerelatesubDeleteAPIResponse {
	return poolTaobaoPlaceStorerelatesubDeleteAPIResponse.Get().(*TaobaoPlaceStorerelatesubDeleteAPIResponse)
}

// ReleaseTaobaoPlaceStorerelatesubDeleteAPIResponse 将 TaobaoPlaceStorerelatesubDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStorerelatesubDeleteAPIResponse(v *TaobaoPlaceStorerelatesubDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStorerelatesubDeleteAPIResponse.Put(v)
}
