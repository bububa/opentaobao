package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取最近报表生成时间 APIResponse
alibaba.scbp.effect.account.date.get

获取最近报表生成时间,格式为yyyy-MM-dd
*/
type AlibabaScbpEffectAccountDateGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpEffectAccountDateGetResponse `json:"alibaba_scbp_effect_account_date_get_response,omitempty"` 
    AlibabaScbpEffectAccountDateGetResponse
}

/* model for simplify = false
type AlibabaScbpEffectAccountDateGetResponse struct {

    // 最近生成报表的时间(US)
    
    ReportDate   string `json:"report_date,omitempty"`
    

}
*/

type AlibabaScbpEffectAccountDateGetResponse struct {

    // 最近生成报表的时间(US)
    ReportDate   string `json:"report_date,omitempty"`

}
