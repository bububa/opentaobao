package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道订单详情查询 API请求
taobao.xhotel.distribution.order.detail.search

该接口用于提供酒店分销渠道的订单详情查询
*/
type TaobaoXhotelDistributionOrderDetailSearchRequest struct {
    model.Params
    // 外部分销渠道的订单号（与tid必填其一）
    _distributionOid   string
    // 传入用户对应的openId
    _openId   string
    // 飞猪的订单号（与distribution_oid必填其一）
    _tid   int64
}

// 初始化TaobaoXhotelDistributionOrderDetailSearchRequest对象
func NewTaobaoXhotelDistributionOrderDetailSearchRequest() *TaobaoXhotelDistributionOrderDetailSearchRequest{
    return &TaobaoXhotelDistributionOrderDetailSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetApiMethodName() string {
    return "taobao.xhotel.distribution.order.detail.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DistributionOid Setter
// 外部分销渠道的订单号（与tid必填其一）
func (r *TaobaoXhotelDistributionOrderDetailSearchRequest) SetDistributionOid(_distributionOid string) error {
    r._distributionOid = _distributionOid
    r.Set("distribution_oid", _distributionOid)
    return nil
}

// DistributionOid Getter
func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetDistributionOid() string {
    return r._distributionOid
}
// OpenId Setter
// 传入用户对应的openId
func (r *TaobaoXhotelDistributionOrderDetailSearchRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetOpenId() string {
    return r._openId
}
// Tid Setter
// 飞猪的订单号（与distribution_oid必填其一）
func (r *TaobaoXhotelDistributionOrderDetailSearchRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetTid() int64 {
    return r._tid
}
