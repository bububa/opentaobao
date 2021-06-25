package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广单元小时级别实时报表查询 APIResponse
taobao.simba.hour.report.adgroup.get

推广单元小时级别实时报表查询
*/
type TaobaoSimbaHourReportAdgroupGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaHourReportAdgroupGetResponse `json:"taobao_simba_hour_report_adgroup_get_response,omitempty"`
}

type TaobaoSimbaHourReportAdgroupGetResponse struct {

    // 11
    Results   *RtRptResultEntityDTO `json:"results,omitempty"`

}
