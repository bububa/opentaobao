package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据健康ID获取支付宝ID APIResponse
alibaba.alihealth.docbase.userinfo.alipayid.get

根据健康ID获取支付宝ID
*/
type AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDocbaseUserinfoAlipayidGetResponse
}

type AlibabaAlihealthDocbaseUserinfoAlipayidGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_docbase_userinfo_alipayid_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
