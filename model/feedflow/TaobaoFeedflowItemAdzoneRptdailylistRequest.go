package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资源包分日数据查询 API请求
taobao.feedflow.item.adzone.rptdailylist

资源包分日数据查询
*/
type TaobaoFeedflowItemAdzoneRptdailylistRequest struct {
    model.Params
    // 查询参数
    rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowItemAdzoneRptdailylistRequest对象
func NewTaobaoFeedflowItemAdzoneRptdailylistRequest() *TaobaoFeedflowItemAdzoneRptdailylistRequest{
    return &TaobaoFeedflowItemAdzoneRptdailylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adzone.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询参数
func (r *TaobaoFeedflowItemAdzoneRptdailylistRequest) SetRptQueryDTO(rptQueryDTO *RptQueryDto) error {
    r.rptQueryDTO = rptQueryDTO
    r.Set("rpt_query_d_t_o", rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemAdzoneRptdailylistRequest) GetRptQueryDTO() *RptQueryDto {
    return r.rptQueryDTO
}
