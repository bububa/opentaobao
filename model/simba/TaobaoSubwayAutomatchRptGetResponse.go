package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询流量智选天级报告 APIResponse
taobao.subway.automatch.rpt.get

查询流量智选天级报告
*/
type TaobaoSubwayAutomatchRptGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSubwayAutomatchRptGetResponse `json:"taobao_subway_automatch_rpt_get_response,omitempty"`
}

type TaobaoSubwayAutomatchRptGetResponse struct {

    // 流量智选天级别报表数据
    ResultList   []ResultMap `json:"result_list,omitempty"`

}
