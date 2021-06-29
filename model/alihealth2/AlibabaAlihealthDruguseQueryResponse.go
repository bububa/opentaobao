package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
合理用药规则查询 APIResponse
alibaba.alihealth.druguse.query

查询用户购买的药品命中的风险规则
*/
type AlibabaAlihealthDruguseQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDruguseQueryResponse
}

type AlibabaAlihealthDruguseQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_druguse_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
