package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔菜单信息上报 APIResponse
taobao.jst.sms.menuinfo.report

聚石塔菜单信息上报
*/
type TaobaoJstSmsMenuinfoReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_sms_menuinfo_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统异常
    
    ResponseCode   string `json:"response_code,omitempty" xml:"