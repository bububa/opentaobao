package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流状态同步 API返回值 
alibaba.ascp.industry.logistics.sync

履约物流状态同步
*/
type AlibabaAscpIndustryLogisticsSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAscpIndustryLogisticsSyncAPIResponseModel
}

// 物流状态同步 成功返回结果
type AlibabaAscpIndustryLogisticsSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ascp_industry_logistics_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *AlibabaAscpIndustryLogisticsSyncResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
