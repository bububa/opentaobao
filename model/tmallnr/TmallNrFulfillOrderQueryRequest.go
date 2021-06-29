package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售商获取品牌商的单笔订单 APIRequest
tmall.nr.fulfill.order.query

零售商获取品牌商的单笔订单，后端服务有零售商和品牌商的绑定关系，存在开关控制；返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
*/
type TmallNrFulfillOrderQueryRequest struct {
    model.Params

    // 业务标识，dss标识定时送业务；jsd表示极速达业务
    bizIdentity   string 

    // 交易主订单号
    orderId   int64 

    // 预留-扩展信息
    extParam   string 

}

func NewTmallNrFulfillOrderQueryRequest() *TmallNrFulfillOrderQueryRequest{
    return &TmallNrFulfillOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrFulfillOrderQueryRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.order.query"
}

func (r TmallNrFulfillOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrFulfillOrderQueryRequest) SetBizIdentity(bizIdentity string) error {
    r.bizIdentity = bizIdentity
    r.Set("biz_identity", bizIdentity)
    return nil
}

func (r TmallNrFulfillOrderQueryRequest) GetBizIdentity() string {
    return r.bizIdentity
}

func (r *TmallNrFulfillOrderQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TmallNrFulfillOrderQueryRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TmallNrFulfillOrderQueryRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

func (r TmallNrFulfillOrderQueryRequest) GetExtParam() string {
    return r.extParam
}

