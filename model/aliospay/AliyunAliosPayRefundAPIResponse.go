package aliospay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunaliospayrefundAPIResponse 退款接口 API返回值
// aliyun.alios.pay.refund
//
// 商户用来发起退款的接口
type AliyunaliospayrefundAPIResponse struct {
	model.CommonResponse
	AliyunaliospayrefundAPIResponseModel
}

// AliyunaliospayrefundAPIResponseModel is 退款接口 成功返回结果
type AliyunaliospayrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOspayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}
