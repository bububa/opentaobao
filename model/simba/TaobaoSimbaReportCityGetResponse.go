package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取城市维度报表 APIResponse
taobao.simba.report.city.get

获取城市维度报表
*/
type TaobaoSimbaReportCityGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaReportCityGetResponse
}

type TaobaoSimbaReportCityGetResponse struct {
    XMLName xml.Name `xml:"simba_report_city_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 11
    
    Results   *RtRptResultEntityDTO `json:"results,omitempty" xml:"results,omitempty"`

    
}
