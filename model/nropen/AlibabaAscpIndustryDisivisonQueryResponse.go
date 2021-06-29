package nropen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务支持地区列表 APIResponse
alibaba.ascp.industry.disivison.query

商家获取服务支持地区
*/
type AlibabaAscpIndustryDisivisonQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAscpIndustryDisivisonQueryResponse
}

type AlibabaAscpIndustryDisivisonQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_industry_disivison_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Data   *AlibabaAscpIndustryDisivisonQueryData `json:"data,omitempty" xml:"data,omitempty"`

    
}
