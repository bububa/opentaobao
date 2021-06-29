package kclub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-通用知识查询服务 APIResponse
alibaba.kclub.kc.queryknowledge

知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
*/
type AlibabaKclubKcQueryknowledgeAPIResponse struct {
    model.CommonResponse
    AlibabaKclubKcQueryknowledgeResponse
}

type AlibabaKclubKcQueryknowledgeResponse struct {
    XMLName xml.Name `xml:"alibaba_kclub_kc_queryknowledge_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *AlibabaKclubKcQueryknowledgeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
