package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessQueryAPIResponse 订单流水查询接口 API返回值
// taobao.qimen.orderprocess.query
//
// ERP调用订单流水查询接口
type TaobaoQimenOrderprocessQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderprocessQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenOrderprocessQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenOrderprocessQueryAPIResponseModel).Reset()
}

// TaobaoQimenOrderprocessQueryAPIResponseModel is 订单流水查询接口 成功返回结果
type TaobaoQimenOrderprocessQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_orderprocess_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *OrderProcessQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenOrderprocessQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenOrderprocessQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderprocessQueryAPIResponse)
	},
}

// GetTaobaoQimenOrderprocessQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenOrderprocessQueryAPIResponse
func GetTaobaoQimenOrderprocessQueryAPIResponse() *TaobaoQimenOrderprocessQueryAPIResponse {
	return poolTaobaoQimenOrderprocessQueryAPIResponse.Get().(*TaobaoQimenOrderprocessQueryAPIResponse)
}

// ReleaseTaobaoQimenOrderprocessQueryAPIResponse 将 TaobaoQimenOrderprocessQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenOrderprocessQueryAPIResponse(v *TaobaoQimenOrderprocessQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenOrderprocessQueryAPIResponse.Put(v)
}
