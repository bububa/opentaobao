package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动下的优惠信息 API请求
taobao.singletreasure.activity.item.query

分页查询活动下的商品优惠信息
*/
type TaobaoSingletreasureActivityItemQueryAPIRequest struct {
    model.Params
    // 活动Id
    _activityId   int64
    // 页大小
    _pageSize   int64
    // 页码
    _pageNumber   int64
}

// 初始化TaobaoSingletreasureActivityItemQueryAPIRequest对象
func NewTaobaoSingletreasureActivityItemQueryRequest() *TaobaoSingletreasureActivityItemQueryAPIRequest{
    return &TaobaoSingletreasureActivityItemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动Id
func (r *TaobaoSingletreasureActivityItemQueryAPIRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetActivityId() int64 {
    return r._activityId
}
// PageSize Setter
// 页大小
func (r *TaobaoSingletreasureActivityItemQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 页码
func (r *TaobaoSingletreasureActivityItemQueryAPIRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetPageNumber() int64 {
    return r._pageNumber
}
