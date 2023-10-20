package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSingleitemQueryAPIResponse 商品查询接口 API返回值
// taobao.qimen.singleitem.query
//
// 商品查询接口
type TaobaoQimenSingleitemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenSingleitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenSingleitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenSingleitemQueryAPIResponseModel).Reset()
}

// TaobaoQimenSingleitemQueryAPIResponseModel is 商品查询接口 成功返回结果
type TaobaoQimenSingleitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_singleitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenSingleitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenSingleitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSingleitemQueryAPIResponse)
	},
}

// GetTaobaoQimenSingleitemQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenSingleitemQueryAPIResponse
func GetTaobaoQimenSingleitemQueryAPIResponse() *TaobaoQimenSingleitemQueryAPIResponse {
	return poolTaobaoQimenSingleitemQueryAPIResponse.Get().(*TaobaoQimenSingleitemQueryAPIResponse)
}

// ReleaseTaobaoQimenSingleitemQueryAPIResponse 将 TaobaoQimenSingleitemQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenSingleitemQueryAPIResponse(v *TaobaoQimenSingleitemQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenSingleitemQueryAPIResponse.Put(v)
}
