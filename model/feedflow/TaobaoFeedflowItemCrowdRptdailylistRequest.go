package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向分日数据查询 APIRequest
taobao.feedflow.item.crowd.rptdailylist

定向分日数据查询
*/
type TaobaoFeedflowItemCrowdRptdailylistRequest struct {
    model.Params

    // 查询条件
    rptQueryDTO   *RptQueryDto 

}

func NewTaobaoFeedflowItemCrowdRptdailylistRequest() *TaobaoFeedflowItemCrowdRptdailylistRequest{
    return &TaobaoFeedflowItemCrowdRptdailylistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCrowdRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.rptdailylist"
}

func (r TaobaoFeedflowItemCrowdRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCrowdRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

func (r TaobaoFeedflowItemCrowdRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}

