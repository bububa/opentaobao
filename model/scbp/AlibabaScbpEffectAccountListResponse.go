package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
账户-报表 APIResponse
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。
*/
type AlibabaScbpEffectAccountListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_effect_account_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 账户效果数据列表
    
    AccountReportList   []AccountEffectDto `json:"account_report_list,omitempty" xml:"