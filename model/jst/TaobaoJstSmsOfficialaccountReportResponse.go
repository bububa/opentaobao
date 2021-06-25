package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号信息上报 APIResponse
taobao.jst.sms.officialaccount.report

聚石塔公众号信息上报
*/
type TaobaoJstSmsOfficialaccountReportAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstSmsOfficialaccountReportResponse `json:"taobao_jst_sms_officialaccount_report_response,omitempty"`
}

type TaobaoJstSmsOfficialaccountReportResponse struct {

    // 系统异常
    ResponseCode   string `json:"response_code,omitempty"`

    // 成功
    ResponseSuccess   bool `json:"response_success,omitempty"`

    // 请求id
    ResponseId   string `json:"response_id,omitempty"`

    // 上报成功
    Module   bool `json:"module,omitempty"`

    // 系统异常
    Message   string `json:"message,omitempty"`

}
