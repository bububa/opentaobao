package aliospay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款接口 APIResponse
aliyun.alios.pay.refund

商户用来发起退款的接口
*/
type AliyunAliosPayRefundAPIResponse struct {
    model.CommonResponse
    AliyunAliosPayRefundResponse
}

type AliyunAliosPayRefundResponse struct {
    XMLName xml.Name `xml:"aliyun_alios_pay_refund_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应参数
    
    AliospayResponse   *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`

    
}
