package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单退款按ID查询 APIRequest
alibaba.wdk.order.refund.get

按照退款ID或者五道口中台订单ID查询退款信息详情
*/
type AlibabaWdkOrderRefundGetRequest struct {
    model.Params

    // 五道口订单列表（子订单）
    bizOrderIds   []int64 

    // 退款订单列表
    refundIds   []int64 

    // 渠道来源 3：饿了么 4：盒马
    orderFrom   int64 

    // 渠道店id
    shopId   string 

    // 经营店id
    storeId   string 

}

func NewAlibabaWdkOrderRefundGetRequest() *AlibabaWdkOrderRefundGetRequest{
    return &AlibabaWdkOrderRefundGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderRefundGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.refund.get"
}

func (r AlibabaWdkOrderRefundGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderRefundGetRequest) SetBizOrderIds(bizOrderIds []int64) error {
    r.bizOrderIds = bizOrderIds
    r.Set("biz_order_ids", bizOrderIds)
    return nil
}

func (r AlibabaWdkOrderRefundGetRequest) GetBizOrderIds() []int64 {
    return r.bizOrderIds
}

func (r *AlibabaWdkOrderRefundGetRequest) SetRefundIds(refundIds []int64) error {
    r.refundIds = refundIds
    r.Set("refund_ids", refundIds)
    return nil
}

func (r AlibabaWdkOrderRefundGetRequest) GetRefundIds() []int64 {
    return r.refundIds
}

func (r *AlibabaWdkOrderRefundGetRequest) SetOrderFrom(orderFrom int64) error {
    r.orderFrom = orderFrom
    r.Set("order_from", orderFrom)
    return nil
}

func (r AlibabaWdkOrderRefundGetRequest) GetOrderFrom() int64 {
    return r.orderFrom
}

func (r *AlibabaWdkOrderRefundGetRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaWdkOrderRefundGetRequest) GetShopId() string {
    return r.shopId
}

func (r *AlibabaWdkOrderRefundGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkOrderRefundGetRequest) GetStoreId() string {
    return r.storeId
}

