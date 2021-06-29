package foodscan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据scanId查询报告 APIResponse
alibaba.footscan.mini.query.mobilereport

根据scanId查询报告
*/
type AlibabaFootscanMiniQueryMobilereportAPIResponse struct {
    model.CommonResponse
    AlibabaFootscanMiniQueryMobilereportResponse
}

type AlibabaFootscanMiniQueryMobilereportResponse struct {
    XMLName xml.Name `xml:"alibaba_footscan_mini_query_mobilereport_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Result   *AlibabaFootscanMiniQueryMobilereportMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
