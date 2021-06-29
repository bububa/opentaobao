package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元分日数据查询 APIRequest
taobao.feedflow.item.adgroup.rptdailylist

推广单元分日数据查询
*/
type TaobaoFeedflowItemAdgroupRptdailylistRequest struct {
    model.Params

    // 查询条件
    rptQueryDTO   *RptQueryDto 

}

func NewTaobaoFeedflowItemAdgroupRptdailylistRequest() *TaobaoFeedflowItemAdgroupRptdailylistRequest{
    return &TaobaoFeedflowItemAdgroupRptdailylistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.rptdailylist"
}

func (r TaobaoFeedflowItemAdgroupRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

func (r TaobaoFeedflowItemAdgroupRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}

