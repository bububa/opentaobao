package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车查询订单id API请求
tmall.car.lease.orderid.get

天猫开新车查询订单id
*/
type TmallCarLeaseOrderidGetRequest struct {
    model.Params
    // openid
    _openId   string
}

// 初始化TmallCarLeaseOrderidGetRequest对象
func NewTmallCarLeaseOrderidGetRequest() *TmallCarLeaseOrderidGetRequest{
    return &TmallCarLeaseOrderidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseOrderidGetRequest) GetApiMethodName() string {
    return "tmall.car.lease.orderid.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseOrderidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenId Setter
// openid
func (r *TmallCarLeaseOrderidGetRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r TmallCarLeaseOrderidGetRequest) GetOpenId() string {
    return r._openId
}
