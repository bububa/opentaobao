package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔菜单信息上报 APIResponse
taobao.jst.sms.menuinfo.report

聚石塔菜单信息上报
*/
type TaobaoJstSmsMenuinfoReportAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstSmsMenuinfoReportResponse `json:"taobao_jst_sms_menuinfo_report_response,omitempty"`
}

type TaobaoJstSmsMenuinfoReportResponse struct {

    // 系统异常
    ResponseCode   string `json:"response_code,omitempty"`

    // 成功
    ResponseSuccess   bool `json:"response_success,omitempty"`

    // 请求id
    ResponseId   string `json:"response_id,omitempty"`

    // 成功
    Module   bool `json:"module,omitempty"`

    // 系统异常
    ResponseMessage   string `json:"response_message,omitempty"`

}
