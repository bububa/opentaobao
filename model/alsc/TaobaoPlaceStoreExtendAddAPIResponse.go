package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreExtendAddAPIResponse 新增门店扩展属性 API返回值
// taobao.place.store.extend.add
//
// 新增授权用户的门店扩展属性
type TaobaoPlaceStoreExtendAddAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreExtendAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreExtendAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreExtendAddAPIResponseModel).Reset()
}

// TaobaoPlaceStoreExtendAddAPIResponseModel is 新增门店扩展属性 成功返回结果
type TaobaoPlaceStoreExtendAddAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_extend_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// 返回结果：true成功；false失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreExtendAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TotalNum = 0
	m.Failure = false
	m.Result = false
}

var poolTaobaoPlaceStoreExtendAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreExtendAddAPIResponse)
	},
}

// GetTaobaoPlaceStoreExtendAddAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreExtendAddAPIResponse
func GetTaobaoPlaceStoreExtendAddAPIResponse() *TaobaoPlaceStoreExtendAddAPIResponse {
	return poolTaobaoPlaceStoreExtendAddAPIResponse.Get().(*TaobaoPlaceStoreExtendAddAPIResponse)
}

// ReleaseTaobaoPlaceStoreExtendAddAPIResponse 将 TaobaoPlaceStoreExtendAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreExtendAddAPIResponse(v *TaobaoPlaceStoreExtendAddAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreExtendAddAPIResponse.Put(v)
}
