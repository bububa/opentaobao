package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后合同下载 API请求
tmall.car.lease.contractdownload

天猫开新车租后合同下载
*/
type TmallCarLeaseContractdownloadRequest struct {
    model.Params
    // 天猫开新车订单id
    _orderId   int64
    // 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
    _type   string
}

// 初始化TmallCarLeaseContractdownloadRequest对象
func NewTmallCarLeaseContractdownloadRequest() *TmallCarLeaseContractdownloadRequest{
    return &TmallCarLeaseContractdownloadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseContractdownloadRequest) GetApiMethodName() string {
    return "tmall.car.lease.contractdownload"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseContractdownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeaseContractdownloadRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeaseContractdownloadRequest) GetOrderId() int64 {
    return r._orderId
}
// Type Setter
// 续租协议： 1， 全款购车协议： 2，分期购买协议：3， 分期购买车辆资产验收协议：4，分期购买车辆抵押：5， 分期购买融资租赁合同：6
func (r *TmallCarLeaseContractdownloadRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TmallCarLeaseContractdownloadRequest) GetType() string {
    return r._type
}