package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取广告主分日数据 API请求
taobao.feedflow.account.rptdailylist

获取广告主分日数据
*/
type TaobaoFeedflowAccountRptdailylistRequest struct {
    model.Params
    // 查询条件
    _rptQueryDTO   *RptQueryDTO
}

// 初始化TaobaoFeedflowAccountRptdailylistRequest对象
func NewTaobaoFeedflowAccountRptdailylistRequest() *TaobaoFeedflowAccountRptdailylistRequest{
    return &TaobaoFeedflowAccountRptdailylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowAccountRptdailylistRequest) GetApiMethodName() string {
    return "taobao.feedflow.account.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowAccountRptdailylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowAccountRptdailylistRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDTO) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowAccountRptdailylistRequest) GetRptQueryDTO() *RptQueryDTO {
    return r._rptQueryDTO
}
