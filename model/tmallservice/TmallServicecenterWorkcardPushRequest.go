package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送服务工单信息 API请求
tmall.servicecenter.workcard.push

服务商家推送工单信息到天猫。
*/
type TmallServicecenterWorkcardPushRequest struct {
    model.Params
    // 属性列表。使用半角分号隔开,字符串前后都需要有半角分号
    _attributes   string
    // 描述
    _desc   string
    // 淘宝交易订单号
    _bizOrderId   int64
    // 服务预约安装时间
    _serviceReserveTime   string
    // 服务预约安装地址。四级地址与街道地址用空格隔开
    _serviceReserveAddress   string
    // 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
    _status   string
}

// 初始化TmallServicecenterWorkcardPushRequest对象
func NewTmallServicecenterWorkcardPushRequest() *TmallServicecenterWorkcardPushRequest{
    return &TmallServicecenterWorkcardPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardPushRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.push"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Attributes Setter
// 属性列表。使用半角分号隔开,字符串前后都需要有半角分号
func (r *TmallServicecenterWorkcardPushRequest) SetAttributes(_attributes string) error {
    r._attributes = _attributes
    r.Set("attributes", _attributes)
    return nil
}

// Attributes Getter
func (r TmallServicecenterWorkcardPushRequest) GetAttributes() string {
    return r._attributes
}
// Desc Setter
// 描述
func (r *TmallServicecenterWorkcardPushRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TmallServicecenterWorkcardPushRequest) GetDesc() string {
    return r._desc
}
// BizOrderId Setter
// 淘宝交易订单号
func (r *TmallServicecenterWorkcardPushRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TmallServicecenterWorkcardPushRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
// ServiceReserveTime Setter
// 服务预约安装时间
func (r *TmallServicecenterWorkcardPushRequest) SetServiceReserveTime(_serviceReserveTime string) error {
    r._serviceReserveTime = _serviceReserveTime
    r.Set("service_reserve_time", _serviceReserveTime)
    return nil
}

// ServiceReserveTime Getter
func (r TmallServicecenterWorkcardPushRequest) GetServiceReserveTime() string {
    return r._serviceReserveTime
}
// ServiceReserveAddress Setter
// 服务预约安装地址。四级地址与街道地址用空格隔开
func (r *TmallServicecenterWorkcardPushRequest) SetServiceReserveAddress(_serviceReserveAddress string) error {
    r._serviceReserveAddress = _serviceReserveAddress
    r.Set("service_reserve_address", _serviceReserveAddress)
    return nil
}

// ServiceReserveAddress Getter
func (r TmallServicecenterWorkcardPushRequest) GetServiceReserveAddress() string {
    return r._serviceReserveAddress
}
// Status Setter
// 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
func (r *TmallServicecenterWorkcardPushRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TmallServicecenterWorkcardPushRequest) GetStatus() string {
    return r._status
}
