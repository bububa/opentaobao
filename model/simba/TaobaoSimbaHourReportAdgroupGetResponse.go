package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元小时级别实时报表查询 APIResponse
taobao.simba.hour.report.adgroup.get

推广单元小时级别实时报表查询
*/
type TaobaoSimbaHourReportAdgroupGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaHourReportAdgroupGetResponse
}

type TaobaoSimbaHourReportAdgroupGetResponse struct {
    XMLName xml.Name `xml:"simba_hour_report_adgroup_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 11
    
    Results   *RtRptResultEntityDTO `json:"results,omitempty" xml:"results,omitempty"`

    
}
