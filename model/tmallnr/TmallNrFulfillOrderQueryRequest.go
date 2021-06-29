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
type TmallNrFulfillOrderQueryRequest struct {
    model.Params
    // 业务标识，dss标识定时送业务；jsd表示极速达业务
    _bizIdentity   string
    // 交易主订单号
    _orderId   int64
    // 预留-扩展信息
    _extParam   string
}

// 初始化TmallNrFulfillOrderQueryRequest对象
func NewTmallNrFulfillOrderQueryRequest() *TmallNrFulfillOrderQueryRequest{
    return &TmallNrFulfillOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrFulfillOrderQueryRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.order.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrFulfillOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallNrFulfillOrderQueryRequest) SetBizIdentity(_bizIdentity string) error {
    r._bizIdentity = _bizIdentity
    r.Set("biz_identity", _bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrFulfillOrderQueryRequest) GetBizIdentity() string {
    return r._bizIdentity
}
// OrderId Setter
// 交易主订单号
func (r *TmallNrFulfillOrderQueryRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallNrFulfillOrderQueryRequest) GetOrderId() int64 {
    return r._orderId
}
// ExtParam Setter
// 预留-扩展信息
func (r *TmallNrFulfillOrderQueryRequest) SetExtParam(_extParam string) error {
    r._extParam = _extParam
    r.Set("ext_param", _extParam)
    return nil
}

// ExtParam Getter
func (r TmallNrFulfillOrderQueryRequest) GetExtParam() string {
    return r._extParam
}
