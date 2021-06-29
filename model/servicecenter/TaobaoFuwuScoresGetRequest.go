package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台评价查询接口 API请求
taobao.fuwu.scores.get

根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
*/
type TaobaoFuwuScoresGetRequest struct {
    model.Params
    // 当前页
    _currentPage   int64
    // 每页获取条数。默认值40，最小值1，最大值100。
    _pageSize   int64
    // 评价日期，查询某一天的评价
    _date   string
}

// 初始化TaobaoFuwuScoresGetRequest对象
func NewTaobaoFuwuScoresGetRequest() *TaobaoFuwuScoresGetRequest{
    return &TaobaoFuwuScoresGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFuwuScoresGetRequest) GetApiMethodName() string {
    return "taobao.fuwu.scores.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFuwuScoresGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentPage Setter
// 当前页
func (r *TaobaoFuwuScoresGetRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoFuwuScoresGetRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值100。
func (r *TaobaoFuwuScoresGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFuwuScoresGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// Date Setter
// 评价日期，查询某一天的评价
func (r *TaobaoFuwuScoresGetRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r TaobaoFuwuScoresGetRequest) GetDate() string {
    return r._date
}
