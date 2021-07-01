package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询汽车租赁活动信息 API请求
tmall.car.lease.item.activity.get

查询汽车租赁活动信息
*/
type TmallCarLeaseItemActivityGetAPIRequest struct {
    model.Params
}

// 初始化TmallCarLeaseItemActivityGetAPIRequest对象
func NewTmallCarLeaseItemActivityGetRequest() *TmallCarLeaseItemActivityGetAPIRequest{
    return &TmallCarLeaseItemActivityGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseItemActivityGetAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.item.activity.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseItemActivityGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
