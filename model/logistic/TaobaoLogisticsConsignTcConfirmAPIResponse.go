package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsConsignTcConfirmAPIResponse 通知交易确认发货接口 API返回值
// taobao.logistics.consign.tc.confirm
//
// 下述业务场景可以使用此接口通知相关的交易订单发货：
// 1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
// 2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。
type TaobaoLogisticsConsignTcConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsConsignTcConfirmAPIResponseModel
}

// TaobaoLogisticsConsignTcConfirmAPIResponseModel is 通知交易确认发货接口 成功返回结果
type TaobaoLogisticsConsignTcConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_consign_tc_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 菜鸟发货单据
	OrderConsignCode string `json:"order_consign_code,omitempty" xml:"order_consign_code,omitempty"`
	// 单次调用流程唯一id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 是否重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
