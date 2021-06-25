package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询今日消耗 APIResponse
alibaba.scbp.account.daycost.get

查询今日消耗
*/
type AlibabaScbpAccountDaycostGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAccountDaycostGetResponse `json:"alibaba_scbp_account_daycost_get_response,omitempty"`
}

type AlibabaScbpAccountDaycostGetResponse struct {

    // 返回今日消耗，单位元，两位小数
    DayCost   string `json:"day_cost,omitempty"`

}
