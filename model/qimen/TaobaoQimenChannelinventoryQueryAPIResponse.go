package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenChannelinventoryQueryAPIResponse 渠道库存查询接口 API返回值
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
type TaobaoQimenChannelinventoryQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenChannelinventoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenChannelinventoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenChannelinventoryQueryAPIResponseModel).Reset()
}

// TaobaoQimenChannelinventoryQueryAPIResponseModel is 渠道库存查询接口 成功返回结果
type TaobaoQimenChannelinventoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_channelinventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenChannelinventoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenChannelinventoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenChannelinventoryQueryAPIResponse)
	},
}

// GetTaobaoQimenChannelinventoryQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenChannelinventoryQueryAPIResponse
func GetTaobaoQimenChannelinventoryQueryAPIResponse() *TaobaoQimenChannelinventoryQueryAPIResponse {
	return poolTaobaoQimenChannelinventoryQueryAPIResponse.Get().(*TaobaoQimenChannelinventoryQueryAPIResponse)
}

// ReleaseTaobaoQimenChannelinventoryQueryAPIResponse 将 TaobaoQimenChannelinventoryQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenChannelinventoryQueryAPIResponse(v *TaobaoQimenChannelinventoryQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenChannelinventoryQueryAPIResponse.Put(v)
}
