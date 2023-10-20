package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoregroupDeleteAPIResponse 删除门店库 API返回值
// taobao.place.storegroup.delete
//
// 删除门店库
type TaobaoPlaceStoregroupDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoregroupDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoregroupDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoregroupDeleteAPIResponseModel).Reset()
}

// TaobaoPlaceStoregroupDeleteAPIResponseModel is 删除门店库 成功返回结果
type TaobaoPlaceStoregroupDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storegroup_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoregroupDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStoregroupDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoregroupDeleteAPIResponse)
	},
}

// GetTaobaoPlaceStoregroupDeleteAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoregroupDeleteAPIResponse
func GetTaobaoPlaceStoregroupDeleteAPIResponse() *TaobaoPlaceStoregroupDeleteAPIResponse {
	return poolTaobaoPlaceStoregroupDeleteAPIResponse.Get().(*TaobaoPlaceStoregroupDeleteAPIResponse)
}

// ReleaseTaobaoPlaceStoregroupDeleteAPIResponse 将 TaobaoPlaceStoregroupDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoregroupDeleteAPIResponse(v *TaobaoPlaceStoregroupDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoregroupDeleteAPIResponse.Put(v)
}
