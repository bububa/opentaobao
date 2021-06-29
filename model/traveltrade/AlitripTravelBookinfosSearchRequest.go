package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单预定信息列表搜索接口 APIRequest
alitrip.travel.bookinfos.search

查询订单预定信息列表
*/
type AlitripTravelBookinfosSearchRequest struct {
    model.Params

    // 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
    pageSize   int64 

    // 当前页
    currentPage   int64 

    // 申请时间_结束，精确到分钟
    applyTimeEnd   string 

    // 申请时间_开始，精确到分钟
    applyTimeStart   string 

}

func NewAlitripTravelBookinfosSearchRequest() *AlitripTravelBookinfosSearchRequest{
    return &AlitripTravelBookinfosSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelBookinfosSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.bookinfos.search"
}

func (r AlitripTravelBookinfosSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelBookinfosSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlitripTravelBookinfosSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlitripTravelBookinfosSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlitripTravelBookinfosSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlitripTravelBookinfosSearchRequest) SetApplyTimeEnd(applyTimeEnd string) error {
    r.applyTimeEnd = applyTimeEnd
    r.Set("apply_time_end", applyTimeEnd)
    return nil
}

func (r AlitripTravelBookinfosSearchRequest) GetApplyTimeEnd() string {
    return r.applyTimeEnd
}

func (r *AlitripTravelBookinfosSearchRequest) SetApplyTimeStart(applyTimeStart string) error {
    r.applyTimeStart = applyTimeStart
    r.Set("apply_time_start", applyTimeStart)
    return nil
}

func (r AlitripTravelBookinfosSearchRequest) GetApplyTimeStart() string {
    return r.applyTimeStart
}

