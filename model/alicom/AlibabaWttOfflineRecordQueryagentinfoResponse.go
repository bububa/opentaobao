package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下推广充送等业务订单来源 APIResponse
alibaba.wtt.offline.record.queryagentinfo

线下推广充送等业务订单来源的查询接口
*/
type AlibabaWttOfflineRecordQueryagentinfoAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWttOfflineRecordQueryagentinfoResponse `json:"alibaba_wtt_offline_record_queryagentinfo_response,omitempty"` 
    AlibabaWttOfflineRecordQueryagentinfoResponse
}

/* model for simplify = false
type AlibabaWttOfflineRecordQueryagentinfoResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWttOfflineRecordQueryagentinfoResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
