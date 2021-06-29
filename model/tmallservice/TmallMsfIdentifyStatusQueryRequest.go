package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅定案核销状态查询 API请求
tmall.msf.identify.status.query

喵师傅定案核销状态查询，供服务商erp系统调用
*/
type TmallMsfIdentifyStatusQueryRequest struct {
    model.Params
    // 天猫订单号
    _orderId   int64
    // 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
    _serviceType   int64
}

// 初始化TmallMsfIdentifyStatusQueryRequest对象
func NewTmallMsfIdentifyStatusQueryRequest() *TmallMsfIdentifyStatusQueryRequest{
    return &TmallMsfIdentifyStatusQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMsfIdentifyStatusQueryRequest) GetApiMethodName() string {
    return "tmall.msf.identify.status.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallMsfIdentifyStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫订单号
func (r *TmallMsfIdentifyStatusQueryRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallMsfIdentifyStatusQueryRequest) GetOrderId() int64 {
    return r._orderId
}
// ServiceType Setter
// 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
func (r *TmallMsfIdentifyStatusQueryRequest) SetServiceType(_serviceType int64) error {
    r._serviceType = _serviceType
    r.Set("service_type", _serviceType)
    return nil
}

// ServiceType Getter
func (r TmallMsfIdentifyStatusQueryRequest) GetServiceType() int64 {
    return r._serviceType
}
