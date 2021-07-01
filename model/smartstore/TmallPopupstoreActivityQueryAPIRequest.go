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
type TmallPopupstoreActivityQueryAPIRequest struct {
    model.Params
    // 查询开始时间,yyyy-MM-dd
    _startDate   string
    // 查询结束时间，yyyy-MM-dd
    _endDate   string
}

// 初始化TmallPopupstoreActivityQueryAPIRequest对象
func NewTmallPopupstoreActivityQueryRequest() *TmallPopupstoreActivityQueryAPIRequest{
    return &TmallPopupstoreActivityQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPopupstoreActivityQueryAPIRequest) GetApiMethodName() string {
    return "tmall.popupstore.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallPopupstoreActivityQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 查询开始时间,yyyy-MM-dd
func (r *TmallPopupstoreActivityQueryAPIRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TmallPopupstoreActivityQueryAPIRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 查询结束时间，yyyy-MM-dd
func (r *TmallPopupstoreActivityQueryAPIRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TmallPopupstoreActivityQueryAPIRequest) GetEndDate() string {
    return r._endDate
}
