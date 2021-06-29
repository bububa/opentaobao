package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流状态同步 APIResponse
alibaba.ascp.industry.logistics.sync

履约物流状态同步
*/
type AlibabaAscpIndustryLogisticsSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAscpIndustryLogisticsSyncResponse
}

type AlibabaAscpIndustryLogisticsSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_industry_logistics_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaAscpIndustryLogisticsSyncResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
