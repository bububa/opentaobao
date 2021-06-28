package kclub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询租户下类目树 APIResponse
alibaba.kclub.kc.getcategorytree

知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
*/
type AlibabaKclubKcGetcategorytreeAPIResponse struct {
    model.CommonResponse
    AlibabaKclubKcGetcategorytreeResponse
}

type AlibabaKclubKcGetcategorytreeResponse struct {
    XMLName xml.Name `xml:"alibaba_kclub_kc_getcategorytree_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaKclubKcGetcategorytreeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
