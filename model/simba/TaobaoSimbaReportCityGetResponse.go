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
    Response *TaobaoSimbaReportCityGetResponse `json:"taobao_simba_report_city_get_response,omitempty"`
}

type TaobaoSimbaReportCityGetResponse struct {

    // 11
    Results   *RtRptResultEntityDTO `json:"results,omitempty"`

}
