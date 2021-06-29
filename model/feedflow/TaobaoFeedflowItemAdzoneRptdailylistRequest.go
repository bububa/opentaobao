package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资源包分日数据查询 APIRequest
taobao.feedflow.item.adzone.rptdailylist

资源包分日数据查询
*/
type TaobaoFeedflowItemAdzoneRptdailylistRequest struct {
    model.Params

    // 查询参数
    rptQueryDTO   *RptQueryDto 

}

func NewTaobaoFeedflowItemAdzoneRptdailylistRequest() *TaobaoFeedflowItemAdzoneRptdailylistRequest{
    return &TaobaoFeedflowItemAdzoneRptdailylistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdzoneRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adzone.rptdailylist"
}

func (r TaobaoFeedflowItemAdzoneRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdzoneRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

func (r TaobaoFeedflowItemAdzoneRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}

