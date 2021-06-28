package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下推广充送等业务订单来源 APIResponse
alibaba.wtt.offline.record.queryagentinfo

线下推广充送等业务订单来源的查询接口
*/
type AlibabaWttOfflineRecordQueryagentinfoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wtt_offline_record_queryagentinfo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"