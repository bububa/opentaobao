package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某段时间内的快闪活动列表 APIRequest
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

func NewTmallPopupstoreActivityQueryRequest() *TmallPopupstoreActivityQueryRequest{
    return &TmallPopupstoreActivityQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPopupstoreActivityQueryRequest) GetApiMethodName() string {
    return "tmall.popupstore.activity.query"
}

func (r TmallPopupstoreActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPopupstoreActivityQueryRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TmallPopupstoreActivityQueryRequest) GetStartDate() string {
    return r.startDate
}

func (r *TmallPopupstoreActivityQueryRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TmallPopupstoreActivityQueryRequest) GetEndDate() string {
    return r.endDate
}

