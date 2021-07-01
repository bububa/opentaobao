package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardBindcardAPIResponse
绑定物理卡 API返回值
alibaba.alsc.crm.card.bindcard

将会员卡和实例物理卡绑定在一起 */
type AlibabaAlscCrmCardBindcardAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardBindcardAPIResponseModel
}

// AlibabaAlscCrmCardBindcardAPIResponseModel is 绑定物理卡 成功返回结果
type AlibabaAlscCrmCardBindcardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_bindcard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
