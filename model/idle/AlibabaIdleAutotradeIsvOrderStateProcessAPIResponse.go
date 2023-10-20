package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleautotradeisvorderstateprocessAPIResponse 闲鱼订单状态推进 API返回值
// alibaba.idle.autotrade.isv.order.state.process
//
// 闲鱼订单状态推进
type AlibabaidleautotradeisvorderstateprocessAPIResponse struct {
	model.CommonResponse
	AlibabaidleautotradeisvorderstateprocessAPIResponseModel
}

// AlibabaidleautotradeisvorderstateprocessAPIResponseModel is 闲鱼订单状态推进 成功返回结果
type AlibabaidleautotradeisvorderstateprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_autotrade_isv_order_state_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
