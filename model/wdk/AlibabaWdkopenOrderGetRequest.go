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
    _storeId   string
    // 五道口主订单id
    _bizOrderId   int64
    // 外部主订单ID
    _outOrderId   string
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
func (r *AlibabaWdkopenOrderGetRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkopenOrderGetRequest) GetStoreId() string {
    return r._storeId
}
// BizOrderId Setter
// 五道口主订单id
func (r *AlibabaWdkopenOrderGetRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaWdkopenOrderGetRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
// OutOrderId Setter
// 外部主订单ID
func (r *AlibabaWdkopenOrderGetRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaWdkopenOrderGetRequest) GetOutOrderId() string {
    return r._outOrderId
}
