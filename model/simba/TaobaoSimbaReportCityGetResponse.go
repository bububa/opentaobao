package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取城市维度报表 APIResponse
taobao.simba.report.city.get

获取城市维度报表
*/
type TaobaoSimbaReportCityGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaReportCityGetResponse `json:"simba_report_city_get_response,omitempty"` 
    TaobaoSimbaReportCityGetResponse
}

/* model for simplify = false
type TaobaoSimbaReportCityGetResponse struct {

    // 11
    
    Results  *struct {
        RtRptResultEntityDTO  *RtRptResultEntityDTO `json:"rt_rpt_result_entity_dto,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoSimbaReportCityGetResponse struct {

    // 11
    Results   *RtRptResultEntityDTO `json:"results,omitempty"`

}
