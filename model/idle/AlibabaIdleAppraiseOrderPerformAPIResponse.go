package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleappraiseorderperformAPIResponse 闲鱼验货宝服务商订单履约 API返回值
// alibaba.idle.appraise.order.perform
//
// 闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
type AlibabaidleappraiseorderperformAPIResponse struct {
	model.CommonResponse
	AlibabaidleappraiseorderperformAPIResponseModel
}

// AlibabaidleappraiseorderperformAPIResponseModel is 闲鱼验货宝服务商订单履约 成功返回结果
type AlibabaidleappraiseorderperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_appraise_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaidleappraiseorderperformResult `json:"result,omitempty" xml:"result,omitempty"`
}
