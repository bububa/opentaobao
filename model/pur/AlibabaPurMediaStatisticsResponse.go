package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新媒体统计信息 APIResponse
alibaba.pur.media.statistics

清博同步新媒体的统计信息给到采购平台
*/
type AlibabaPurMediaStatisticsAPIResponse struct {
    model.CommonResponse
    AlibabaPurMediaStatisticsResponse
}

type AlibabaPurMediaStatisticsResponse struct {
    XMLName xml.Name `xml:"alibaba_pur_media_statistics_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取url的出参
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
