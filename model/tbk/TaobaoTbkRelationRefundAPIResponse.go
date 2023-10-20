package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkrelationrefundAPIResponse 淘宝客-推广者-维权退款订单查询 API返回值
// taobao.tbk.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
type TaobaotbkrelationrefundAPIResponse struct {
	model.CommonResponse
	TaobaotbkrelationrefundAPIResponseModel
}

// TaobaotbkrelationrefundAPIResponseModel is 淘宝客-推广者-维权退款订单查询 成功返回结果
type TaobaotbkrelationrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_relation_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果封装对象
	Result *TaobaotbkrelationrefundRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}
