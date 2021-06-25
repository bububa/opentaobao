package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户发起逆向取货 APIRequest
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

func NewAlibabaTclsAelophyRefundFetchgoodsRequest() *AlibabaTclsAelophyRefundFetchgoodsRequest{
    return &AlibabaTclsAelophyRefundFetchgoodsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.fetchgoods"
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetOutOrderId() string {
    return r.outOrderId
}

func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetRefundId() string {
    return r.refundId
}

func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetFetchStartTime(fetchStartTime string) error {
    r.fetchStartTime = fetchStartTime
    r.Set("fetch_start_time", fetchStartTime)
    return nil
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetFetchStartTime() string {
    return r.fetchStartTime
}

func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetFetchEndTime(fetchEndTime string) error {
    r.fetchEndTime = fetchEndTime
    r.Set("fetch_end_time", fetchEndTime)
    return nil
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetFetchEndTime() string {
    return r.fetchEndTime
}

func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetRemark() string {
    return r.remark
}

func (r *AlibabaTclsAelophyRefundFetchgoodsRequest) SetSubRefundList(subRefundList []Subrefundlist) error {
    r.subRefundList = subRefundList
    r.Set("sub_refund_list", subRefundList)
    return nil
}

func (r AlibabaTclsAelophyRefundFetchgoodsRequest) GetSubRefundList() []Subrefundlist {
    return r.subRefundList
}

