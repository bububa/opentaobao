package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOnlineSendAPIResponse 在线订单发货处理（支持货到付款） API返回值
// taobao.logistics.online.send
//
// 用户调用该接口可实现在线订单发货（支持货到付款）<br/>调用该接口实现在线下单发货，有两种情况：<br><br/><font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br><br/>如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font>
type TaobaoLogisticsOnlineSendAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOnlineSendAPIResponseModel
}

// TaobaoLogisticsOnlineSendAPIResponseModel is 在线订单发货处理（支持货到付款） 成功返回结果
type TaobaoLogisticsOnlineSendAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_online_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// de
	Shipping *Shipping `json:"shipping,omitempty" xml:"shipping,omitempty"`
}
