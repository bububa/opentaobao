package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemmappingQueryAPIResponse 前后端商品映射查询接口 API返回值
// taobao.qimen.itemmapping.query
//
// 前后端商品映射查询接口
type TaobaoQimenItemmappingQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemmappingQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenItemmappingQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenItemmappingQueryAPIResponseModel).Reset()
}

// TaobaoQimenItemmappingQueryAPIResponseModel is 前后端商品映射查询接口 成功返回结果
type TaobaoQimenItemmappingQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_itemmapping_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenItemmappingQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenItemmappingQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenItemmappingQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemmappingQueryAPIResponse)
	},
}

// GetTaobaoQimenItemmappingQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenItemmappingQueryAPIResponse
func GetTaobaoQimenItemmappingQueryAPIResponse() *TaobaoQimenItemmappingQueryAPIResponse {
	return poolTaobaoQimenItemmappingQueryAPIResponse.Get().(*TaobaoQimenItemmappingQueryAPIResponse)
}

// ReleaseTaobaoQimenItemmappingQueryAPIResponse 将 TaobaoQimenItemmappingQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenItemmappingQueryAPIResponse(v *TaobaoQimenItemmappingQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenItemmappingQueryAPIResponse.Put(v)
}
