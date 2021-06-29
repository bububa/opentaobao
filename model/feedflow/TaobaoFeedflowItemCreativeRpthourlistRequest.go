package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】创意分时报表查询 API请求
taobao.feedflow.item.creative.rpthourlist

创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
*/
type TaobaoFeedflowItemCreativeRpthourlistRequest struct {
    model.Params
    // 查询参数
    rptQuery   *RptQueryDto
}

// 初始化TaobaoFeedflowItemCreativeRpthourlistRequest对象
func NewTaobaoFeedflowItemCreativeRpthourlistRequest() *TaobaoFeedflowItemCreativeRpthourlistRequest{
    return &TaobaoFeedflowItemCreativeRpthourlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowItemCreativeRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowItemCreativeRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}
