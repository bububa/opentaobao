package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsTaowaiMsgSendAPIResponse 淘外包裹物流信息走淘宝发包裹状态通知接口 API返回值
// taobao.logistics.taowai.msg.send
//
// 淘外包裹物流信息走淘宝发包裹状态通知接口
type TaobaoLogisticsTaowaiMsgSendAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsTaowaiMsgSendAPIResponseModel
}

// TaobaoLogisticsTaowaiMsgSendAPIResponseModel is 淘外包裹物流信息走淘宝发包裹状态通知接口 成功返回结果
type TaobaoLogisticsTaowaiMsgSendAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_taowai_msg_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 对象
	Result *MsgSendResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 通讯成功/通讯失败
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
