package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单二次预约查询接口 APIRequest
alitrip.travel.bookinfo.query

飞猪度假订单二次预约详情查询接口
*/
type AlitripTravelBookinfoQueryRequest struct {
    model.Params

    // 预定信息id
    bookinfoId   int64 

}

func NewAlitripTravelBookinfoQueryRequest() *AlitripTravelBookinfoQueryRequest{
    return &AlitripTravelBookinfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelBookinfoQueryRequest) GetApiMethodName() string {
    return "alitrip.travel.bookinfo.query"
}

func (r AlitripTravelBookinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelBookinfoQueryRequest) SetBookinfoId(bookinfoId int64) error {
    r.bookinfoId = bookinfoId
    r.Set("bookinfo_id", bookinfoId)
    return nil
}

func (r AlitripTravelBookinfoQueryRequest) GetBookinfoId() int64 {
    return r.bookinfoId
}

