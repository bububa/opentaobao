package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 APIRequest
taobao.wlb.order.jzwithins.consign

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
*/
type TaobaoWlbOrderJzwithinsConsignRequest struct {
    model.Params

    // 淘宝交易订单号
    tid   int64 

    // 物流服务商信息
    tmsPartner   *JzPartnerNew 

    // 物流服务商信息
    insPartner   *JzPartnerNew 

    // 家装物流发货参数
    jzConsignArgs   *JzConsignArgsNew 

}

func NewTaobaoWlbOrderJzwithinsConsignRequest() *TaobaoWlbOrderJzwithinsConsignRequest{
    return &TaobaoWlbOrderJzwithinsConsignRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderJzwithinsConsignRequest) GetApiMethodName() string {
    return "taobao.wlb.order.jzwithins.consign"
}

func (r TaobaoWlbOrderJzwithinsConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoWlbOrderJzwithinsConsignRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetTmsPartner(tmsPartner *JzPartnerNew) error {
    r.tmsPartner = tmsPartner
    r.Set("tms_partner", tmsPartner)
    return nil
}

func (r TaobaoWlbOrderJzwithinsConsignRequest) GetTmsPartner() *JzPartnerNew {
    return r.tmsPartner
}

func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetInsPartner(insPartner *JzPartnerNew) error {
    r.insPartner = insPartner
    r.Set("ins_partner", insPartner)
    return nil
}

func (r TaobaoWlbOrderJzwithinsConsignRequest) GetInsPartner() *JzPartnerNew {
    return r.insPartner
}

func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetJzConsignArgs(jzConsignArgs *JzConsignArgsNew) error {
    r.jzConsignArgs = jzConsignArgs
    r.Set("jz_consign_args", jzConsignArgs)
    return nil
}

func (r TaobaoWlbOrderJzwithinsConsignRequest) GetJzConsignArgs() *JzConsignArgsNew {
    return r.jzConsignArgs
}

