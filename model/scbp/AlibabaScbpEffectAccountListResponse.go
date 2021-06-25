package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
账户-报表 APIResponse
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。
*/
type AlibabaScbpEffectAccountListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpEffectAccountListResponse `json:"alibaba_scbp_effect_account_list_response,omitempty"`
}

type AlibabaScbpEffectAccountListResponse struct {

    // 账户效果数据列表
    AccountReportList   []AccountEffectDto `json:"account_report_list,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

}
