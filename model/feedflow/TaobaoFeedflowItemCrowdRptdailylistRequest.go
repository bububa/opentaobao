package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向分日数据查询 API请求
taobao.feedflow.item.crowd.rptdailylist

定向分日数据查询
*/
type TaobaoFeedflowItemCrowdRptdailylistAPIRequest struct {
    model.Params
    // 查询条件
    _rptQueryDTO   *RptQueryDTO
}

// 初始化TaobaoFeedflowItemCrowdRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemCrowdRptdailylistRequest() *TaobaoFeedflowItemCrowdRptdailylistAPIRequest{
    return &TaobaoFeedflowItemCrowdRptdailylistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdRptdailylistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdRptdailylistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCrowdRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDTO) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowItemCrowdRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDTO {
    return r._rptQueryDTO
}
