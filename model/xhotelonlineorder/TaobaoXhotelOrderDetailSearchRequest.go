package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情查询 API请求
taobao.xhotel.order.detail.search

提供订单详情查询
*/
type TaobaoXhotelOrderDetailSearchRequest struct {
    model.Params
    // 外部订单号
    _outOid   string
    // 外部订单号
    _tid   int64
}

// 初始化TaobaoXhotelOrderDetailSearchRequest对象
func NewTaobaoXhotelOrderDetailSearchRequest() *TaobaoXhotelOrderDetailSearchRequest{
    return &TaobaoXhotelOrderDetailSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderDetailSearchRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.detail.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderDetailSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutOid Setter
// 外部订单号
func (r *TaobaoXhotelOrderDetailSearchRequest) SetOutOid(_outOid string) error {
    r._outOid = _outOid
    r.Set("out_oid", _outOid)
    return nil
}

// OutOid Getter
func (r TaobaoXhotelOrderDetailSearchRequest) GetOutOid() string {
    return r._outOid
}
// Tid Setter
// 外部订单号
func (r *TaobaoXhotelOrderDetailSearchRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderDetailSearchRequest) GetTid() int64 {
    return r._tid
}
