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
type TmallNrOrderQueryJstRequest struct {
    model.Params
    // 业务标识，dss标识定时送业务；jsd表示极速达业务
    bizIdentity   string
    // 交易主订单号
    orderId   int64
    // 预留-扩展信息
    extParam   string
}

// 初始化TmallNrOrderQueryJstRequest对象
func NewTmallNrOrderQueryJstRequest() *TmallNrOrderQueryJstRequest{
    return &TmallNrOrderQueryJstRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrOrderQueryJstRequest) GetApiMethodName() string {
    return "tmall.nr.order.query.jst"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrOrderQueryJstRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallNrOrderQueryJstRequest) SetBizIdentity(bizIdentity string) error {
    r.bizIdentity = bizIdentity
    r.Set("biz_identity", bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrOrderQueryJstRequest) GetBizIdentity() string {
    return r.bizIdentity
}
// OrderId Setter
// 交易主订单号
func (r *TmallNrOrderQueryJstRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallNrOrderQueryJstRequest) GetOrderId() int64 {
    return r.orderId
}
// ExtParam Setter
// 预留-扩展信息
func (r *TmallNrOrderQueryJstRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

// ExtParam Getter
func (r TmallNrOrderQueryJstRequest) GetExtParam() string {
    return r.extParam
}
