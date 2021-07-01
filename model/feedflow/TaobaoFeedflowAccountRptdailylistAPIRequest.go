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
type TaobaoFeedflowAccountRptdailylistAPIRequest struct {
    model.Params
    // 查询条件
    _rptQueryDTO   *RptQueryDto
}

// 初始化TaobaoFeedflowAccountRptdailylistAPIRequest对象
func NewTaobaoFeedflowAccountRptdailylistRequest() *TaobaoFeedflowAccountRptdailylistAPIRequest{
    return &TaobaoFeedflowAccountRptdailylistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowAccountRptdailylistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.account.rptdailylist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowAccountRptdailylistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowAccountRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
    r._rptQueryDTO = _rptQueryDTO
    r.Set("rpt_query_d_t_o", _rptQueryDTO)
    return nil
}

// RptQueryDTO Getter
func (r TaobaoFeedflowAccountRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
    return r._rptQueryDTO
}
