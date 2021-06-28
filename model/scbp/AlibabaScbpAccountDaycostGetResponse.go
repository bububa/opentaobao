package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询今日消耗 APIResponse
alibaba.scbp.account.daycost.get

查询今日消耗
*/
type AlibabaScbpAccountDaycostGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_account_daycost_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回今日消耗，单位元，两位小数
    
    DayCost   string `json:"day_cost,omitempty" xml:"