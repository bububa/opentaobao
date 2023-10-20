package itpolicy

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareAddrtAPIResponse 【国际机票自有政策】单条往返添加 API返回值
// taobao.alitrip.it.fare.addrt
//
// 自有政策往返添加接口
type TaobaoAlitripItFareAddrtAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItFareAddrtAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareAddrtAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripItFareAddrtAPIResponseModel).Reset()
}

// TaobaoAlitripItFareAddrtAPIResponseModel is 【国际机票自有政策】单条往返添加 成功返回结果
type TaobaoAlitripItFareAddrtAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_addrt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 运价id
	FareId int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareAddrtAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ExtendAttributes = ""
	m.FareId = 0
}

var poolTaobaoAlitripItFareAddrtAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripItFareAddrtAPIResponse)
	},
}

// GetTaobaoAlitripItFareAddrtAPIResponse 从 sync.Pool 获取 TaobaoAlitripItFareAddrtAPIResponse
func GetTaobaoAlitripItFareAddrtAPIResponse() *TaobaoAlitripItFareAddrtAPIResponse {
	return poolTaobaoAlitripItFareAddrtAPIResponse.Get().(*TaobaoAlitripItFareAddrtAPIResponse)
}

// ReleaseTaobaoAlitripItFareAddrtAPIResponse 将 TaobaoAlitripItFareAddrtAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripItFareAddrtAPIResponse(v *TaobaoAlitripItFareAddrtAPIResponse) {
	v.Reset()
	poolTaobaoAlitripItFareAddrtAPIResponse.Put(v)
}
