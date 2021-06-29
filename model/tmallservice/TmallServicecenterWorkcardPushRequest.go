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
    attributes   string
    // 描述
    desc   string
    // 淘宝交易订单号
    bizOrderId   int64
    // 服务预约安装时间
    serviceReserveTime   string
    // 服务预约安装地址。四级地址与街道地址用空格隔开
    serviceReserveAddress   string
    // 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
    status   string
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
func (r *TmallServicecenterWorkcardPushRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

// Attributes Getter
func (r TmallServicecenterWorkcardPushRequest) GetAttributes() string {
    return r.attributes
}
// Desc Setter
// 描述
func (r *TmallServicecenterWorkcardPushRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TmallServicecenterWorkcardPushRequest) GetDesc() string {
    return r.desc
}
// BizOrderId Setter
// 淘宝交易订单号
func (r *TmallServicecenterWorkcardPushRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TmallServicecenterWorkcardPushRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
// ServiceReserveTime Setter
// 服务预约安装时间
func (r *TmallServicecenterWorkcardPushRequest) SetServiceReserveTime(serviceReserveTime string) error {
    r.serviceReserveTime = serviceReserveTime
    r.Set("service_reserve_time", serviceReserveTime)
    return nil
}

// ServiceReserveTime Getter
func (r TmallServicecenterWorkcardPushRequest) GetServiceReserveTime() string {
    return r.serviceReserveTime
}
// ServiceReserveAddress Setter
// 服务预约安装地址。四级地址与街道地址用空格隔开
func (r *TmallServicecenterWorkcardPushRequest) SetServiceReserveAddress(serviceReserveAddress string) error {
    r.serviceReserveAddress = serviceReserveAddress
    r.Set("service_reserve_address", serviceReserveAddress)
    return nil
}

// ServiceReserveAddress Getter
func (r TmallServicecenterWorkcardPushRequest) GetServiceReserveAddress() string {
    return r.serviceReserveAddress
}
// Status Setter
// 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
func (r *TmallServicecenterWorkcardPushRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TmallServicecenterWorkcardPushRequest) GetStatus() string {
    return r.status
}
