package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据健康ID获取支付宝ID API返回值 
alibaba.alihealth.docbase.userinfo.alipayid.get

根据健康ID获取支付宝ID
*/
type AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDocbaseUserinfoAlipayidGetResponse
}

// 根据健康ID获取支付宝ID 成功返回结果
type AlibabaAlihealthDocbaseUserinfoAlipayidGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_docbase_userinfo_alipayid_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
