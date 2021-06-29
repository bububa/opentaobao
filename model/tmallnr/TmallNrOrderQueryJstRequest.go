package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取同城配送业务单笔订单 APIRequest
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

func NewTmallNrOrderQueryJstRequest() *TmallNrOrderQueryJstRequest{
    return &TmallNrOrderQueryJstRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrOrderQueryJstRequest) GetApiMethodName() string {
    return "tmall.nr.order.query.jst"
}

func (r TmallNrOrderQueryJstRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrOrderQueryJstRequest) SetBizIdentity(bizIdentity string) error {
    r.bizIdentity = bizIdentity
    r.Set("biz_identity", bizIdentity)
    return nil
}

func (r TmallNrOrderQueryJstRequest) GetBizIdentity() string {
    return r.bizIdentity
}

func (r *TmallNrOrderQueryJstRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TmallNrOrderQueryJstRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TmallNrOrderQueryJstRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

func (r TmallNrOrderQueryJstRequest) GetExtParam() string {
    return r.extParam
}

