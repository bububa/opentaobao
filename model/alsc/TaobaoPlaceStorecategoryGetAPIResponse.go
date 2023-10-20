package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorecategoryGetAPIResponse 获取门店类目信息 API返回值
// taobao.place.storecategory.get
//
// 获取门店类目信息
type TaobaoPlaceStorecategoryGetAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStorecategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStorecategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStorecategoryGetAPIResponseModel).Reset()
}

// TaobaoPlaceStorecategoryGetAPIResponseModel is 获取门店类目信息 成功返回结果
type TaobaoPlaceStorecategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storecategory_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店类目格式
	CategoryList string `json:"category_list,omitempty" xml:"category_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStorecategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CategoryList = ""
}

var poolTaobaoPlaceStorecategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStorecategoryGetAPIResponse)
	},
}

// GetTaobaoPlaceStorecategoryGetAPIResponse 从 sync.Pool 获取 TaobaoPlaceStorecategoryGetAPIResponse
func GetTaobaoPlaceStorecategoryGetAPIResponse() *TaobaoPlaceStorecategoryGetAPIResponse {
	return poolTaobaoPlaceStorecategoryGetAPIResponse.Get().(*TaobaoPlaceStorecategoryGetAPIResponse)
}

// ReleaseTaobaoPlaceStorecategoryGetAPIResponse 将 TaobaoPlaceStorecategoryGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStorecategoryGetAPIResponse(v *TaobaoPlaceStorecategoryGetAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStorecategoryGetAPIResponse.Put(v)
}
