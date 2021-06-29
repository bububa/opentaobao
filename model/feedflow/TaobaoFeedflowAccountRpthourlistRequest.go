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
type TaobaoFeedflowAccountRpthourlistRequest struct {
    model.Params
    // 查询参数
    rptQuery   *RptQueryDto
}

// 初始化TaobaoFeedflowAccountRpthourlistRequest对象
func NewTaobaoFeedflowAccountRpthourlistRequest() *TaobaoFeedflowAccountRpthourlistRequest{
    return &TaobaoFeedflowAccountRpthourlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowAccountRpthourlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.account.rpthourlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowAccountRpthourlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RptQuery Setter
// 查询参数
func (r *TaobaoFeedflowAccountRpthourlistRequest) SetRptQuery(rptQuery *RptQueryDto) error {
    r.rptQuery = rptQuery
    r.Set("rpt_query", rptQuery)
    return nil
}

// RptQuery Getter
func (r TaobaoFeedflowAccountRpthourlistRequest) GetRptQuery() *RptQueryDto {
    return r.rptQuery
}
