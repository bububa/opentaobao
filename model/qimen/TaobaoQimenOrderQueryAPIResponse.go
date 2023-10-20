package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderQueryAPIResponse 根据收件人信息查询交易单号接口 API返回值
// taobao.qimen.order.query
//
// WMS 调用该接口，根据收件人信息查询平台交易订单号。
type TaobaoQimenOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenOrderQueryAPIResponseModel).Reset()
}

// TaobaoQimenOrderQueryAPIResponseModel is 根据收件人信息查询交易单号接口 成功返回结果
type TaobaoQimenOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Response *TaobaoQimenOrderQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderQueryAPIResponse)
	},
}

// GetTaobaoQimenOrderQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenOrderQueryAPIResponse
func GetTaobaoQimenOrderQueryAPIResponse() *TaobaoQimenOrderQueryAPIResponse {
	return poolTaobaoQimenOrderQueryAPIResponse.Get().(*TaobaoQimenOrderQueryAPIResponse)
}

// ReleaseTaobaoQimenOrderQueryAPIResponse 将 TaobaoQimenOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenOrderQueryAPIResponse(v *TaobaoQimenOrderQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenOrderQueryAPIResponse.Put(v)
}
