package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
统计adgroup数量 APIResponse
alibaba.scbp.ad.group.count.ad.group

统计adgroup数量
*/
type AlibabaScbpAdGroupCountAdGroupAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupCountAdGroupResponse
}

type AlibabaScbpAdGroupCountAdGroupResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_count_ad_group_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
