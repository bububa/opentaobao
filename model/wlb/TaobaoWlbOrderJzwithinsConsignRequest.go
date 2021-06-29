package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 API请求
taobao.wlb.order.jzwithins.consign

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
*/
type TaobaoWlbOrderJzwithinsConsignRequest struct {
    model.Params
    // 淘宝交易订单号
    _tid   int64
    // 物流服务商信息
    _tmsPartner   *JzPartnerNew
    // 物流服务商信息
    _insPartner   *JzPartnerNew
    // 家装物流发货参数
    _jzConsignArgs   *JzConsignArgsNew
}

// 初始化TaobaoWlbOrderJzwithinsConsignRequest对象
func NewTaobaoWlbOrderJzwithinsConsignRequest() *TaobaoWlbOrderJzwithinsConsignRequest{
    return &TaobaoWlbOrderJzwithinsConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderJzwithinsConsignRequest) GetApiMethodName() string {
    return "taobao.wlb.order.jzwithins.consign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderJzwithinsConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易订单号
func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoWlbOrderJzwithinsConsignRequest) GetTid() int64 {
    return r._tid
}
// TmsPartner Setter
// 物流服务商信息
func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetTmsPartner(_tmsPartner *JzPartnerNew) error {
    r._tmsPartner = _tmsPartner
    r.Set("tms_partner", _tmsPartner)
    return nil
}

// TmsPartner Getter
func (r TaobaoWlbOrderJzwithinsConsignRequest) GetTmsPartner() *JzPartnerNew {
    return r._tmsPartner
}
// InsPartner Setter
// 物流服务商信息
func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetInsPartner(_insPartner *JzPartnerNew) error {
    r._insPartner = _insPartner
    r.Set("ins_partner", _insPartner)
    return nil
}

// InsPartner Getter
func (r TaobaoWlbOrderJzwithinsConsignRequest) GetInsPartner() *JzPartnerNew {
    return r._insPartner
}
// JzConsignArgs Setter
// 家装物流发货参数
func (r *TaobaoWlbOrderJzwithinsConsignRequest) SetJzConsignArgs(_jzConsignArgs *JzConsignArgsNew) error {
    r._jzConsignArgs = _jzConsignArgs
    r.Set("jz_consign_args", _jzConsignArgs)
    return nil
}

// JzConsignArgs Getter
func (r TaobaoWlbOrderJzwithinsConsignRequest) GetJzConsignArgs() *JzConsignArgsNew {
    return r._jzConsignArgs
}
