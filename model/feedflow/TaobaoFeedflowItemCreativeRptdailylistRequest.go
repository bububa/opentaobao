package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创意分日数据查询 APIRequest
taobao.feedflow.item.creative.rptdailylist

创意分日数据查询
*/
type TaobaoFeedflowItemCreativeRptdailylistRequest struct {
    model.Params

    // 查询条件
    rptQueryDTO   *RptQueryDto 

}

func NewTaobaoFeedflowItemCreativeRptdailylistRequest() *TaobaoFeedflowItemCreativeRptdailylistRequest{
    return &TaobaoFeedflowItemCreativeRptdailylistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCreativeRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.rptdailylist"
}

func (r TaobaoFeedflowItemCreativeRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCreativeRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

func (r TaobaoFeedflowItemCreativeRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}

