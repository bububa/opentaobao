package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动下的优惠信息 APIRequest
taobao.singletreasure.activity.item.query

分页查询活动下的商品优惠信息
*/
type TaobaoSingletreasureActivityItemQueryRequest struct {
    model.Params

    // 活动Id
    activityId   int64 

    // 页大小
    pageSize   int64 

    // 页码
    pageNumber   int64 

}

func NewTaobaoSingletreasureActivityItemQueryRequest() *TaobaoSingletreasureActivityItemQueryRequest{
    return &TaobaoSingletreasureActivityItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityItemQueryRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.query"
}

func (r TaobaoSingletreasureActivityItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityItemQueryRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoSingletreasureActivityItemQueryRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoSingletreasureActivityItemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSingletreasureActivityItemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSingletreasureActivityItemQueryRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

func (r TaobaoSingletreasureActivityItemQueryRequest) GetPageNumber() int64 {
    return r.pageNumber
}

