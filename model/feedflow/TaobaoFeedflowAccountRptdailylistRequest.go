package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取广告主分日数据 APIRequest
taobao.feedflow.account.rptdailylist

获取广告主分日数据
*/
type TaobaoFeedflowAccountRptdailylistRequest struct {
    model.Params

    // 查询条件
    rptQueryDTO   *RptQueryDto 

}

func NewTaobaoFeedflowAccountRptdailylistRequest() *TaobaoFeedflowAccountRptdailylistRequest{
    return &TaobaoFeedflowAccountRptdailylistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowAccountRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.account.rptdailylist"
}

func (r TaobaoFeedflowAccountRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowAccountRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

func (r TaobaoFeedflowAccountRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}

