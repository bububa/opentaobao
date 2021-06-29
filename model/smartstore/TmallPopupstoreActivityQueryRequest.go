package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某段时间内的快闪活动列表 API请求
tmall.popupstore.activity.query

提供给ISV查询某一时间段内包含指定appKey的活动列表
*/
type TmallPopupstoreActivityQueryRequest struct {
    model.Params
    // 查询开始时间,yyyy-MM-dd
    startDate   string
    // 查询结束时间，yyyy-MM-dd
    endDate   string
}

// 初始化TmallPopupstoreActivityQueryRequest对象
func NewTmallPopupstoreActivityQueryRequest() *TmallPopupstoreActivityQueryRequest{
    return &TmallPopupstoreActivityQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPopupstoreActivityQueryRequest) GetApiMethodName() string {
    return "tmall.popupstore.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallPopupstoreActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 查询开始时间,yyyy-MM-dd
func (r *TmallPopupstoreActivityQueryRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TmallPopupstoreActivityQueryRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 查询结束时间，yyyy-MM-dd
func (r *TmallPopupstoreActivityQueryRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TmallPopupstoreActivityQueryRequest) GetEndDate() string {
    return r.endDate
}
