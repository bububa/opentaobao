package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeUnchargeUpdateAPIResponse
充值退款 API返回值
alibaba.alsc.crm.recharge.uncharge.update

充值退款 */
type AlibabaAlscCrmRechargeUnchargeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeUnchargeUpdateAPIResponseModel
}

// AlibabaAlscCrmRechargeUnchargeUpdateAPIResponseModel is 充值退款 成功返回结果
type AlibabaAlscCrmRechargeUnchargeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_uncharge_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
