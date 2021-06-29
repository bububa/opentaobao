package aliospay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取支付token APIResponse
aliyun.alios.pay.token.get

商户用来获取支付的授权token
*/
type AliyunAliosPayTokenGetAPIResponse struct {
    model.CommonResponse
    AliyunAliosPayTokenGetResponse
}

type AliyunAliosPayTokenGetResponse struct {
    XMLName xml.Name `xml:"aliyun_alios_pay_token_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应参数
    
    AliospayResponse   *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`

    
}
