package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐广告主分时报表查询 API请求
taobao.feedflow.account.rpthourlist

广告主分时报表查询，支持广告主查询最近90天内某一天的账户维度分时报表数据
*/
type TaobaoFeedflowAccountRpthourlistAPIRequest struct {
    model.Params
    // 查询参数
    _rptQuery   *RptQueryDTO
}

// 初始化TaobaoFeedflowAccountRpthourlistAPIRequest对象
func NewTaobaoFeedflowAccountRpthourlistRequest() *TaobaoFeedflowAccountRpthourlistAPIRequest{
    return &TaobaoFeedflowAccountRpthourlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowAccountRpthourlistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.account.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowAccountRpthourlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowAccountRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDTO) error {
    r._rptQuery = _rptQuery
    r.Set("rpt_query", _rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowAccountRpthourlistAPIRequest) GetRptQuery() *RptQueryDTO {
    return r._rptQuery
}
