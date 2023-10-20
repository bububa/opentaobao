package itpolicy

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareAddowAPIResponse 【国际机票自有政策】单条单程添加 API返回值
// taobao.alitrip.it.fare.addow
//
// 自有政策单程添加接口，重复的老数据会被删除，重复判断规则同excel
type TaobaoAlitripItFareAddowAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItFareAddowAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareAddowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripItFareAddowAPIResponseModel).Reset()
}

// TaobaoAlitripItFareAddowAPIResponseModel is 【国际机票自有政策】单条单程添加 成功返回结果
type TaobaoAlitripItFareAddowAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_addow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 运价id
	FareId int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareAddowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ExtendAttributes = ""
	m.FareId = 0
}

var poolTaobaoAlitripItFareAddowAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripItFareAddowAPIResponse)
	},
}

// GetTaobaoAlitripItFareAddowAPIResponse 从 sync.Pool 获取 TaobaoAlitripItFareAddowAPIResponse
func GetTaobaoAlitripItFareAddowAPIResponse() *TaobaoAlitripItFareAddowAPIResponse {
	return poolTaobaoAlitripItFareAddowAPIResponse.Get().(*TaobaoAlitripItFareAddowAPIResponse)
}

// ReleaseTaobaoAlitripItFareAddowAPIResponse 将 TaobaoAlitripItFareAddowAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripItFareAddowAPIResponse(v *TaobaoAlitripItFareAddowAPIResponse) {
	v.Reset()
	poolTaobaoAlitripItFareAddowAPIResponse.Put(v)
}
