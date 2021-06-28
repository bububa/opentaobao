package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取捡漏词包分时报表数据 APIResponse
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据
*/
type TaobaoSubwayMarshLandRptGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSubwayMarshLandRptGetResponse `json:"subway_marsh_land_rpt_get_response,omitempty"` 
    TaobaoSubwayMarshLandRptGetResponse
}

/* model for simplify = false
type TaobaoSubwayMarshLandRptGetResponse struct {

    // 捡漏词包分时报表数据列表
    
    ResultList  struct {
        RptResult  []RptResult `json:"rpt_result,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type TaobaoSubwayMarshLandRptGetResponse struct {

    // 捡漏词包分时报表数据列表
    ResultList   []RptResult `json:"result_list,omitempty"`

}
