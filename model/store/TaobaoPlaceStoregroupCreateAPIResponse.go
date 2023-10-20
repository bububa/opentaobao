package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoregroupCreateAPIResponse 商户门店库创建接口 API返回值
// taobao.place.storegroup.create
//
// 用于商家创建线下门店库
type TaobaoPlaceStoregroupCreateAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoregroupCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoregroupCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoregroupCreateAPIResponseModel).Reset()
}

// TaobaoPlaceStoregroupCreateAPIResponseModel is 商户门店库创建接口 成功返回结果
type TaobaoPlaceStoregroupCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storegroup_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoregroupCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStoregroupCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoregroupCreateAPIResponse)
	},
}

// GetTaobaoPlaceStoregroupCreateAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoregroupCreateAPIResponse
func GetTaobaoPlaceStoregroupCreateAPIResponse() *TaobaoPlaceStoregroupCreateAPIResponse {
	return poolTaobaoPlaceStoregroupCreateAPIResponse.Get().(*TaobaoPlaceStoregroupCreateAPIResponse)
}

// ReleaseTaobaoPlaceStoregroupCreateAPIResponse 将 TaobaoPlaceStoregroupCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoregroupCreateAPIResponse(v *TaobaoPlaceStoregroupCreateAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoregroupCreateAPIResponse.Put(v)
}
