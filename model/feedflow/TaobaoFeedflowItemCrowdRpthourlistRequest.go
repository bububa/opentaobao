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
type TaobaoFeedflowItemCrowdRpthourlistRequest struct {
    model.Params
    // 查询参数
    _rptQuery   *RptQueryDto
}

// 初始化TaobaoFeedflowItemCrowdRpthourlistRequest对象
func NewTaobaoFeedflowItemCrowdRpthourlistRequest() *TaobaoFeedflowItemCrowdRpthourlistRequest{
    return &TaobaoFeedflowItemCrowdRpthourlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemCrowdRpthourlistRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
    r._rptQuery = _rptQuery
    r.Set("rpt_query", _rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowItemCrowdRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r._rptQuery
}
