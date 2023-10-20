package aliospay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayTokenGetAPIResponse 获取支付token API返回值
// aliyun.alios.pay.token.get
//
// 商户用来获取支付的授权token
type AliyunAliosPayTokenGetAPIResponse struct {
	model.CommonResponse
	AliyunAliosPayTokenGetAPIResponseModel
}

// AliyunAliosPayTokenGetAPIResponseModel is 获取支付token 成功返回结果
type AliyunAliosPayTokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOspayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}
