package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkRelationRefundAPIResponse 淘宝客-推广者-维权退款订单查询 API返回值
// taobao.tbk.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
type TaobaoTbkRelationRefundAPIResponse struct {
	model.CommonResponse
	TaobaoTbkRelationRefundAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkRelationRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkRelationRefundAPIResponseModel).Reset()
}

// TaobaoTbkRelationRefundAPIResponseModel is 淘宝客-推广者-维权退款订单查询 成功返回结果
type TaobaoTbkRelationRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_relation_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果封装对象
	Result *TaobaoTbkRelationRefundRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkRelationRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTbkRelationRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkRelationRefundAPIResponse)
	},
}

// GetTaobaoTbkRelationRefundAPIResponse 从 sync.Pool 获取 TaobaoTbkRelationRefundAPIResponse
func GetTaobaoTbkRelationRefundAPIResponse() *TaobaoTbkRelationRefundAPIResponse {
	return poolTaobaoTbkRelationRefundAPIResponse.Get().(*TaobaoTbkRelationRefundAPIResponse)
}

// ReleaseTaobaoTbkRelationRefundAPIResponse 将 TaobaoTbkRelationRefundAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkRelationRefundAPIResponse(v *TaobaoTbkRelationRefundAPIResponse) {
	v.Reset()
	poolTaobaoTbkRelationRefundAPIResponse.Put(v)
}
