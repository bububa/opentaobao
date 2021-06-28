package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
账户级别小时报表获取 APIResponse
taobao.simba.hour.report.account.get

获取账户小时实时报表数据
*/
type TaobaoSimbaHourReportAccountGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaHourReportAccountGetResponse
}

type TaobaoSimbaHourReportAccountGetResponse struct {
    XMLName xml.Name `xml:"simba_hour_report_account_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 11
    
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
    
    
}
