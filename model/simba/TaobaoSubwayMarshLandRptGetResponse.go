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
    Response *TaobaoSubwayMarshLandRptGetResponse `json:"taobao_subway_marsh_land_rpt_get_response,omitempty"`
}

type TaobaoSubwayMarshLandRptGetResponse struct {

    // 捡漏词包分时报表数据列表
    ResultList   []RptResult `json:"result_list,omitempty"`

}
