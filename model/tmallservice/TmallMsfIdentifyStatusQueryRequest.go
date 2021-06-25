package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅定案核销状态查询 APIRequest
tmall.msf.identify.status.query

喵师傅定案核销状态查询，供服务商erp系统调用
*/
type TmallMsfIdentifyStatusQueryRequest struct {
    model.Params

    // 天猫订单号
    orderId   int64 

    // 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
    serviceType   int64 

}

func NewTmallMsfIdentifyStatusQueryRequest() *TmallMsfIdentifyStatusQueryRequest{
    return &TmallMsfIdentifyStatusQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMsfIdentifyStatusQueryRequest) GetApiMethodName() string {
    return "tmall.msf.identify.status.query"
}

func (r TmallMsfIdentifyStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMsfIdentifyStatusQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TmallMsfIdentifyStatusQueryRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TmallMsfIdentifyStatusQueryRequest) SetServiceType(serviceType int64) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

func (r TmallMsfIdentifyStatusQueryRequest) GetServiceType() int64 {
    return r.serviceType
}

