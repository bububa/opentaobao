package itpolicy

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareDeleteAPIResponse 【国际机票自有政策】单条删除 API返回值
// taobao.alitrip.it.fare.delete
//
// 自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败
type TaobaoAlitripItFareDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItFareDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripItFareDeleteAPIResponseModel).Reset()
}

// TaobaoAlitripItFareDeleteAPIResponseModel is 【国际机票自有政策】单条删除 成功返回结果
type TaobaoAlitripItFareDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ExtendAttributes = ""
}

var poolTaobaoAlitripItFareDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripItFareDeleteAPIResponse)
	},
}

// GetTaobaoAlitripItFareDeleteAPIResponse 从 sync.Pool 获取 TaobaoAlitripItFareDeleteAPIResponse
func GetTaobaoAlitripItFareDeleteAPIResponse() *TaobaoAlitripItFareDeleteAPIResponse {
	return poolTaobaoAlitripItFareDeleteAPIResponse.Get().(*TaobaoAlitripItFareDeleteAPIResponse)
}

// ReleaseTaobaoAlitripItFareDeleteAPIResponse 将 TaobaoAlitripItFareDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripItFareDeleteAPIResponse(v *TaobaoAlitripItFareDeleteAPIResponse) {
	v.Reset()
	poolTaobaoAlitripItFareDeleteAPIResponse.Put(v)
}
