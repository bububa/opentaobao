package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询流量智选天级报告 APIResponse
taobao.subway.automatch.rpt.get

查询流量智选天级报告
*/
type TaobaoSubwayAutomatchRptGetAPIResponse struct {
    model.CommonResponse
    TaobaoSubwayAutomatchRptGetResponse
}

type TaobaoSubwayAutomatchRptGetResponse struct {
    XMLName xml.Name `xml:"subway_automatch_rpt_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 流量智选天级别报表数据
    
    ResultList   []ResultMap `json:"result_list,omitempty" xml:"result_list>result_map,omitempty"`
    
    
}
