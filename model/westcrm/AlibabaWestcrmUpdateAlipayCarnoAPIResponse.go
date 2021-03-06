package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmUpdateAlipayCarnoAPIResponse 更新支付宝业务卡号 API返回值
// alibaba.westcrm.update.alipay.carno
//
// 更新支付宝业务卡号
type AlibabaWestcrmUpdateAlipayCarnoAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmUpdateAlipayCarnoAPIResponseModel
}

// AlibabaWestcrmUpdateAlipayCarnoAPIResponseModel is 更新支付宝业务卡号 成功返回结果
type AlibabaWestcrmUpdateAlipayCarnoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_update_alipay_carno_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回修改结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
