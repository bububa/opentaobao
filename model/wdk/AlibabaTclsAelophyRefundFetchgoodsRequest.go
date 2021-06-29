package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户发起逆向取货 API请求
alibaba.tcls.aelophy.refund.fetchgoods

saas 售后逆向 商户发起逆向取货
*/
type AlibabaTclsAelophyRefundFetchgoodsRequest struct {
    model.Params
    // 经营店ID
    storeId   string
    // 外部订单ID
    outOrderId   string
    // 退款单ID
    refundId   string
    // 取货开始时间
    fetchStartTime   string
    // 取货结束时间
    fetchEndTime   string
    // 备注
    remark   string
    // 外部子订单列表
    subRefundList   []Subrefundlist
}

// 初始化AlibabaTclsAelophyRefundFetchgoodsRequest对象
func NewAlibabaTclsAelophyRefundFetchgoodsRequest() *AlibabaTclsAelophyRefundFetchgoodsRequest{
    return &AlibabaTclsAelophyRefundFetchgoodsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.fetchgoods"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 经营店ID
func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetStoreId() string {
    return r.storeId
}
// OutOrderId Setter
// 外部订单ID
func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetOutOrderId() string {
    return r.outOrderId
}
// RefundId Setter
// 退款单ID
func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetRefundId() string {
    return r.refundId
}
// FetchStartTime Setter
// 取货开始时间
func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetFetchStartTime(fetchStartTime string) error {
    r.fetchStartTime = fetchStartTime
    r.Set("fetch_start_time", fetchStartTime)
    return nil
}

// FetchStartTime Getter
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetFetchStartTime() string {
    return r.fetchStartTime
}
// FetchEndTime Setter
// 取货结束时间
func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetFetchEndTime(fetchEndTime string) error {
    r.fetchEndTime = fetchEndTime
    r.Set("fetch_end_time", fetchEndTime)
    return nil
}

// FetchEndTime Getter
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetFetchEndTime() string {
    return r.fetchEndTime
}
// Remark Setter
// 备注
func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetRemark() string {
    return r.remark
}
// SubRefundList Setter
// 外部子订单列表
func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetSubRefundList(subRefundList []Subrefundlist) error {
    r.subRefundList = subRefundList
    r.Set("sub_refund_list", subRefundList)
    return nil
}

// SubRefundList Getter
func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetSubRefundList() []Subrefundlist {
    return r.subRefundList
}
