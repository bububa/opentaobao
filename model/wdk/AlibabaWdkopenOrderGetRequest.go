package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商户订单获取 APIRequest
alibaba.wdkopen.order.get

商户通过五道口订单id获取订单信息
*/
type AlibabaWdkopenOrderGetRequest struct {
    model.Params

    // 经营店id
    storeId   string 

    // 五道口主订单id
    bizOrderId   int64 

    // 外部主订单ID
    outOrderId   string 

}

func NewAlibabaWdkopenOrderGetRequest() *AlibabaWdkopenOrderGetRequest{
    return &AlibabaWdkopenOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkopenOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdkopen.order.get"
}

func (r AlibabaWdkopenOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkopenOrderGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkopenOrderGetRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkopenOrderGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaWdkopenOrderGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

func (r *AlibabaWdkopenOrderGetRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

func (r AlibabaWdkopenOrderGetRequest) GetOutOrderId() string {
    return r.outOrderId
}

