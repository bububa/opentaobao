package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemmappingCreateAPIResponse 前后端商品映射接口 API返回值
// taobao.qimen.itemmapping.create
//
// 前后端商品映射
type TaobaoQimenItemmappingCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemmappingCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenItemmappingCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenItemmappingCreateAPIResponseModel).Reset()
}

// TaobaoQimenItemmappingCreateAPIResponseModel is 前后端商品映射接口 成功返回结果
type TaobaoQimenItemmappingCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_itemmapping_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenItemmappingCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenItemmappingCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemmappingCreateAPIResponse)
	},
}

// GetTaobaoQimenItemmappingCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenItemmappingCreateAPIResponse
func GetTaobaoQimenItemmappingCreateAPIResponse() *TaobaoQimenItemmappingCreateAPIResponse {
	return poolTaobaoQimenItemmappingCreateAPIResponse.Get().(*TaobaoQimenItemmappingCreateAPIResponse)
}

// ReleaseTaobaoQimenItemmappingCreateAPIResponse 将 TaobaoQimenItemmappingCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenItemmappingCreateAPIResponse(v *TaobaoQimenItemmappingCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenItemmappingCreateAPIResponse.Put(v)
}
