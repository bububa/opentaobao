package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预售结算数据回传 API请求
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。
*/
type AlibabaMjPresaleSettlementAddlistRequest struct {
    model.Params
    // 订单json格式数据
    _preSaleRefundJson   string
}

// 初始化AlibabaMjPresaleSettlementAddlistRequest对象
func NewAlibabaMjPresaleSettlementAddlistRequest() *AlibabaMjPresaleSettlementAddlistRequest{
    return &AlibabaMjPresaleSettlementAddlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjPresaleSettlementAddlistRequest) GetApiMethodName() string {
    return "alibaba.mj.presale.settlement.addlist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjPresaleSettlementAddlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PreSaleRefundJson Setter
// 订单json格式数据
func (r *AlibabaMjPresaleSettlementAddlistRequest) SetPreSaleRefundJson(_preSaleRefundJson string) error {
    r._preSaleRefundJson = _preSaleRefundJson
    r.Set("pre_sale_refund_json", _preSaleRefundJson)
    return nil
}

// PreSaleRefundJson Getter
func (r AlibabaMjPresaleSettlementAddlistRequest) GetPreSaleRefundJson() string {
    return r._preSaleRefundJson
}
