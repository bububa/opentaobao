package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspayregisterauditAPIResponse 阿信支付入驻审核通知 API返回值
// taobao.alitrip.axin.trans.pay.register.audit
//
// 阿信支付入驻审核通知
type TaobaoalitripaxintranspayregisterauditAPIResponse struct {
	model.CommonResponse
	TaobaoalitripaxintranspayregisterauditAPIResponseModel
}

// TaobaoalitripaxintranspayregisterauditAPIResponseModel is 阿信支付入驻审核通知 成功返回结果
type TaobaoalitripaxintranspayregisterauditAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_register_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
