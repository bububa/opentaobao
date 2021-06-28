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
    // Response *TaobaoSimbaHourReportAdgroupGetResponse `json:"simba_hour_report_adgroup_get_response,omitempty"` 
    TaobaoSimbaHourReportAdgroupGetResponse
}

/* model for simplify = false
type TaobaoSimbaHourReportAdgroupGetResponse struct {

    // 11
    
    Results  *struct {
        RtRptResultEntityDTO  *RtRptResultEntityDTO `json:"rt_rpt_result_entity_dto,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoSimbaHourReportAdgroupGetResponse struct {

    // 11
    Results   *RtRptResultEntityDTO `json:"results,omitempty"`

}
