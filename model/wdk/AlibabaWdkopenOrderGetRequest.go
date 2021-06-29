package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商户订单获取 API请求
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

// 初始化AlibabaWdkopenOrderGetRequest对象
func NewAlibabaWdkopenOrderGetRequest() *AlibabaWdkopenOrderGetRequest{
    return &AlibabaWdkopenOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkopenOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdkopen.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkopenOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 经营店id
func (r *AlibabaWdkopenOrderGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkopenOrderGetRequest) GetStoreId() string {
    return r.storeId
}
// BizOrderId Setter
// 五道口主订单id
func (r *AlibabaWdkopenOrderGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaWdkopenOrderGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
// OutOrderId Setter
// 外部主订单ID
func (r *AlibabaWdkopenOrderGetRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaWdkopenOrderGetRequest) GetOutOrderId() string {
    return r.outOrderId
}
