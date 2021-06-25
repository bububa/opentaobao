package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
账户级别小时报表获取 APIResponse
taobao.simba.hour.report.account.get

获取账户小时实时报表数据
*/
type TaobaoSimbaHourReportAccountGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaHourReportAccountGetResponse `json:"taobao_simba_hour_report_account_get_response,omitempty"`
}

type TaobaoSimbaHourReportAccountGetResponse struct {

    // 11
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
