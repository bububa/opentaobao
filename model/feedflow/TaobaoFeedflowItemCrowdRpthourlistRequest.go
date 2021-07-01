package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】定向分时报表查询 API请求
taobao.feedflow.item.crowd.rpthourlist

广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据
*/
type TaobaoFeedflowItemCrowdRpthourlistAPIRequest struct {
    model.Params
    // 查询参数
    _rptQuery   *RptQueryDTO
}

// 初始化TaobaoFeedflowItemCrowdRpthourlistAPIRequest对象
func NewTaobaoFeedflowItemCrowdRpthourlistRequest() *TaobaoFeedflowItemCrowdRpthourlistAPIRequest{
    return &TaobaoFeedflowItemCrowdRpthourlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdRpthourlistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdRpthourlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemCrowdRpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDTO) error {
    r._rptQuery = _rptQuery
    r.Set("rpt_query", _rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowItemCrowdRpthourlistAPIRequest) GetRptQuery() *RptQueryDTO {
    return r._rptQuery
}
