package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台评价查询接口 APIRequest
taobao.fuwu.scores.get

根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
*/
type TaobaoFuwuScoresGetRequest struct {
    model.Params

    // 当前页
    currentPage   int64 

    // 每页获取条数。默认值40，最小值1，最大值100。
    pageSize   int64 

    // 评价日期，查询某一天的评价
    date   string 

}

func NewTaobaoFuwuScoresGetRequest() *TaobaoFuwuScoresGetRequest{
    return &TaobaoFuwuScoresGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFuwuScoresGetRequest) GetApiMethodName() string {
    return "taobao.fuwu.scores.get"
}

func (r TaobaoFuwuScoresGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFuwuScoresGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoFuwuScoresGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TaobaoFuwuScoresGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoFuwuScoresGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoFuwuScoresGetRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r TaobaoFuwuScoresGetRequest) GetDate() string {
    return r.date
}

