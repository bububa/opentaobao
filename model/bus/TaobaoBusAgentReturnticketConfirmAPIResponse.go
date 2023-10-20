package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusagentreturnticketconfirmAPIResponse 商家回调退票 API返回值
// taobao.bus.agent.returnticket.confirm
//
// 商家通过TOP接口调用来回传退票状态
type TaobaobusagentreturnticketconfirmAPIResponse struct {
	model.CommonResponse
	TaobaobusagentreturnticketconfirmAPIResponseModel
}

// TaobaobusagentreturnticketconfirmAPIResponseModel is 商家回调退票 成功返回结果
type TaobaobusagentreturnticketconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_agent_returnticket_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否确认成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
