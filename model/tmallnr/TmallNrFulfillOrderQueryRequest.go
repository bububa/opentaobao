package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售商获取品牌商的单笔订单 API请求
tmall.nr.fulfill.order.query

零售商获取品牌商的单笔订单，后端服务有零售商和品牌商的绑定关系，存在开关控制；返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
*/
type TmallNrFulfillOrderQueryAPIRequest struct {
    model.Params
    // 业务标识，dss标识定时送业务；jsd表示极速达业务
    _bizIdentity   string
    // 交易主订单号
    _orderId   int64
    // 预留-扩展信息
    _extParam   string
}

// 初始化TmallNrFulfillOrderQueryAPIRequest对象
func NewTmallNrFulfillOrderQueryRequest() *TmallNrFulfillOrderQueryAPIRequest{
    return &TmallNrFulfillOrderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrFulfillOrderQueryAPIRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.order.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrFulfillOrderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallNrFulfillOrderQueryAPIRequest) SetBizIdentity(_bizIdentity string) error {
    r._bizIdentity = _bizIdentity
    r.Set("biz_identity", _bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrFulfillOrderQueryAPIRequest) GetBizIdentity() string {
    return r._bizIdentity
}
// OrderId Setter
// 交易主订单号
func (r *TmallNrFulfillOrderQueryAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallNrFulfillOrderQueryAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// ExtParam Setter
// 预留-扩展信息
func (r *TmallNrFulfillOrderQueryAPIRequest) SetExtParam(_extParam string) error {
    r._extParam = _extParam
    r.Set("ext_param", _extParam)
    return nil
}

// ExtParam Getter
func (r TmallNrFulfillOrderQueryAPIRequest) GetExtParam() string {
    return r._extParam
}
