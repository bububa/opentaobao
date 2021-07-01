package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取同城配送业务单笔订单 API请求
tmall.nr.order.query.jst

同城配送业务获取单笔订单
*/
type TmallNrOrderQueryJstAPIRequest struct {
    model.Params
    // 业务标识，dss标识定时送业务；jsd表示极速达业务
    _bizIdentity   string
    // 交易主订单号
    _orderId   int64
    // 预留-扩展信息
    _extParam   string
}

// 初始化TmallNrOrderQueryJstAPIRequest对象
func NewTmallNrOrderQueryJstRequest() *TmallNrOrderQueryJstAPIRequest{
    return &TmallNrOrderQueryJstAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrOrderQueryJstAPIRequest) GetApiMethodName() string {
    return "tmall.nr.order.query.jst"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrOrderQueryJstAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallNrOrderQueryJstAPIRequest) SetBizIdentity(_bizIdentity string) error {
    r._bizIdentity = _bizIdentity
    r.Set("biz_identity", _bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrOrderQueryJstAPIRequest) GetBizIdentity() string {
    return r._bizIdentity
}
// OrderId Setter
// 交易主订单号
func (r *TmallNrOrderQueryJstAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallNrOrderQueryJstAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// ExtParam Setter
// 预留-扩展信息
func (r *TmallNrOrderQueryJstAPIRequest) SetExtParam(_extParam string) error {
    r._extParam = _extParam
    r.Set("ext_param", _extParam)
    return nil
}

// ExtParam Getter
func (r TmallNrOrderQueryJstAPIRequest) GetExtParam() string {
    return r._extParam
}
