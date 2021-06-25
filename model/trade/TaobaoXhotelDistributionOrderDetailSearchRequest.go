package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道订单详情查询 APIRequest
taobao.xhotel.distribution.order.detail.search

该接口用于提供酒店分销渠道的订单详情查询
*/
type TaobaoXhotelDistributionOrderDetailSearchRequest struct {
    model.Params

    // 外部分销渠道的订单号（与tid必填其一）
    distributionOid   string 

    // 传入用户对应的openId
    openId   string 

    // 飞猪的订单号（与distribution_oid必填其一）
    tid   int64 

}

func NewTaobaoXhotelDistributionOrderDetailSearchRequest() *TaobaoXhotelDistributionOrderDetailSearchRequest{
    return &TaobaoXhotelDistributionOrderDetailSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetApiMethodName() string {
    return "taobao.xhotel.distribution.order.detail.search"
}

func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelDistributionOrderDetailSearchRequest) SetDistributionOid(distributionOid string) error {
    r.distributionOid = distributionOid
    r.Set("distribution_oid", distributionOid)
    return nil
}

func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetDistributionOid() string {
    return r.distributionOid
}

func (r *TaobaoXhotelDistributionOrderDetailSearchRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetOpenId() string {
    return r.openId
}

func (r *TaobaoXhotelDistributionOrderDetailSearchRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelDistributionOrderDetailSearchRequest) GetTid() int64 {
    return r.tid
}

