package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenCombineitemQueryAPIResponse 组合货品关系查询接口 API返回值
// taobao.qimen.combineitem.query
//
// 组合货品关系查询
type TaobaoQimenCombineitemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenCombineitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenCombineitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenCombineitemQueryAPIResponseModel).Reset()
}

// TaobaoQimenCombineitemQueryAPIResponseModel is 组合货品关系查询接口 成功返回结果
type TaobaoQimenCombineitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_combineitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenCombineitemQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenCombineitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenCombineitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemQueryAPIResponse)
	},
}

// GetTaobaoQimenCombineitemQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenCombineitemQueryAPIResponse
func GetTaobaoQimenCombineitemQueryAPIResponse() *TaobaoQimenCombineitemQueryAPIResponse {
	return poolTaobaoQimenCombineitemQueryAPIResponse.Get().(*TaobaoQimenCombineitemQueryAPIResponse)
}

// ReleaseTaobaoQimenCombineitemQueryAPIResponse 将 TaobaoQimenCombineitemQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenCombineitemQueryAPIResponse(v *TaobaoQimenCombineitemQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenCombineitemQueryAPIResponse.Put(v)
}
