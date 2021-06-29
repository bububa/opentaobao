package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台工单创建接口 API请求
alibaba.servicecenter.workcard.create

创建服务平台工单
*/
type AlibabaServicecenterWorkcardCreateRequest struct {
    model.Params
    // 服务单id
    spServiceOrderId   int64
    // 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
    serviceSequence   int64
    // 申请次数
    serviceCount   int64
    // 工单属性，json格式字符串
    attributes   string
    // 工单外部唯一键单号
    outerId   string
    // 服务提供者信息
    serviceProvider   *ServiceProviderDto
}

// 初始化AlibabaServicecenterWorkcardCreateRequest对象
func NewAlibabaServicecenterWorkcardCreateRequest() *AlibabaServicecenterWorkcardCreateRequest{
    return &AlibabaServicecenterWorkcardCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardCreateRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.workcard.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterWorkcardCreateRequest) SetSpServiceOrderId(spServiceOrderId int64) error {
    r.spServiceOrderId = spServiceOrderId
    r.Set("sp_service_order_id", spServiceOrderId)
    return nil
}

// SpServiceOrderId Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetSpServiceOrderId() int64 {
    return r.spServiceOrderId
}
// ServiceSequence Setter
// 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
func (r *AlibabaServicecenterWorkcardCreateRequest) SetServiceSequence(serviceSequence int64) error {
    r.serviceSequence = serviceSequence
    r.Set("service_sequence", serviceSequence)
    return nil
}

// ServiceSequence Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetServiceSequence() int64 {
    return r.serviceSequence
}
// ServiceCount Setter
// 申请次数
func (r *AlibabaServicecenterWorkcardCreateRequest) SetServiceCount(serviceCount int64) error {
    r.serviceCount = serviceCount
    r.Set("service_count", serviceCount)
    return nil
}

// ServiceCount Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetServiceCount() int64 {
    return r.serviceCount
}
// Attributes Setter
// 工单属性，json格式字符串
func (r *AlibabaServicecenterWorkcardCreateRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

// Attributes Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetAttributes() string {
    return r.attributes
}
// OuterId Setter
// 工单外部唯一键单号
func (r *AlibabaServicecenterWorkcardCreateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetOuterId() string {
    return r.outerId
}
// ServiceProvider Setter
// 服务提供者信息
func (r *AlibabaServicecenterWorkcardCreateRequest) SetServiceProvider(serviceProvider *ServiceProviderDto) error {
    r.serviceProvider = serviceProvider
    r.Set("service_provider", serviceProvider)
    return nil
}

// ServiceProvider Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetServiceProvider() *ServiceProviderDto {
    return r.serviceProvider
}
