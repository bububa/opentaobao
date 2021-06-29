package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单二次预约查询接口 API请求
alitrip.travel.bookinfo.query

飞猪度假订单二次预约详情查询接口
*/
type AlitripTravelBookinfoQueryRequest struct {
    model.Params
    // 预定信息id
    bookinfoId   int64
}

// 初始化AlitripTravelBookinfoQueryRequest对象
func NewAlitripTravelBookinfoQueryRequest() *AlitripTravelBookinfoQueryRequest{
    return &AlitripTravelBookinfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelBookinfoQueryRequest) GetApiMethodName() string {
    return "alitrip.travel.bookinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelBookinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BookinfoId Setter
// 预定信息id
func (r *AlitripTravelBookinfoQueryRequest) SetBookinfoId(bookinfoId int64) error {
    r.bookinfoId = bookinfoId
    r.Set("bookinfo_id", bookinfoId)
    return nil
}

// BookinfoId Getter
func (r AlitripTravelBookinfoQueryRequest) GetBookinfoId() int64 {
    return r.bookinfoId
}
