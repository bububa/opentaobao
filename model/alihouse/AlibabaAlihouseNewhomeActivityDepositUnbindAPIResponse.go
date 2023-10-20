package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeactivitydepositunbindAPIResponse 销售活动解绑预存金商品 API返回值
// alibaba.alihouse.newhome.activity.deposit.unbind
//
// 销售活动解绑预存金商品
type AlibabaalihousenewhomeactivitydepositunbindAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeactivitydepositunbindAPIResponseModel
}

// AlibabaalihousenewhomeactivitydepositunbindAPIResponseModel is 销售活动解绑预存金商品 成功返回结果
type AlibabaalihousenewhomeactivitydepositunbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_deposit_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文属性
	Result *AlibabaalihousenewhomeactivitydepositunbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
