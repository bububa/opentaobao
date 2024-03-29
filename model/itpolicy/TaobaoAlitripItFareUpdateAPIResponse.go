package itpolicy

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareUpdateAPIResponse 【国际机票自有政策】单条修改 API返回值
// taobao.alitrip.it.fare.update
//
// 自有政策修改接口，可以根据fareId或outId修改，outId不唯一时，不能用outId修改。当外部政策id、出发城市、到达城市、出票航司任一有变化，或往返时是否允许混舱、文件编号、可混文件编号任一有变化，将删除老数据，产生一条新政策。
type TaobaoAlitripItFareUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItFareUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripItFareUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripItFareUpdateAPIResponseModel is 【国际机票自有政策】单条修改 成功返回结果
type TaobaoAlitripItFareUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 运价id，根据更新的内容，此id可能会重新生成
	FareId int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripItFareUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ExtendAttributes = ""
	m.FareId = 0
}

var poolTaobaoAlitripItFareUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripItFareUpdateAPIResponse)
	},
}

// GetTaobaoAlitripItFareUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripItFareUpdateAPIResponse
func GetTaobaoAlitripItFareUpdateAPIResponse() *TaobaoAlitripItFareUpdateAPIResponse {
	return poolTaobaoAlitripItFareUpdateAPIResponse.Get().(*TaobaoAlitripItFareUpdateAPIResponse)
}

// ReleaseTaobaoAlitripItFareUpdateAPIResponse 将 TaobaoAlitripItFareUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripItFareUpdateAPIResponse(v *TaobaoAlitripItFareUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripItFareUpdateAPIResponse.Put(v)
}
