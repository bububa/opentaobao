package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼帮卖同步服务商交易统计信息 APIResponse
alibaba.idle.consignment.spu.statistics

闲鱼帮卖同步服务商交易统计信息
*/
type AlibabaIdleConsignmentSpuStatisticsAPIResponse struct {
    model.CommonResponse
    AlibabaIdleConsignmentSpuStatisticsResponse
}

type AlibabaIdleConsignmentSpuStatisticsResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_consignment_spu_statistics_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
