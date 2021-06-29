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
    _spServiceOrderId   int64
    // 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
    _serviceSequence   int64
    // 申请次数
    _serviceCount   int64
    // 工单属性，json格式字符串
    _attributes   string
    // 工单外部唯一键单号
    _outerId   string
    // 服务提供者信息
    _serviceProvider   *ServiceProviderDto
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
func (r *AlibabaServicecenterWorkcardCreateRequest) SetSpServiceOrderId(_spServiceOrderId int64) error {
    r._spServiceOrderId = _spServiceOrderId
    r.Set("sp_service_order_id", _spServiceOrderId)
    return nil
}

// SpServiceOrderId Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetSpServiceOrderId() int64 {
    return r._spServiceOrderId
}
// ServiceSequence Setter
// 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
func (r *AlibabaServicecenterWorkcardCreateRequest) SetServiceSequence(_serviceSequence int64) error {
    r._serviceSequence = _serviceSequence
    r.Set("service_sequence", _serviceSequence)
    return nil
}

// ServiceSequence Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetServiceSequence() int64 {
    return r._serviceSequence
}
// ServiceCount Setter
// 申请次数
func (r *AlibabaServicecenterWorkcardCreateRequest) SetServiceCount(_serviceCount int64) error {
    r._serviceCount = _serviceCount
    r.Set("service_count", _serviceCount)
    return nil
}

// ServiceCount Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetServiceCount() int64 {
    return r._serviceCount
}
// Attributes Setter
// 工单属性，json格式字符串
func (r *AlibabaServicecenterWorkcardCreateRequest) SetAttributes(_attributes string) error {
    r._attributes = _attributes
    r.Set("attributes", _attributes)
    return nil
}

// Attributes Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetAttributes() string {
    return r._attributes
}
// OuterId Setter
// 工单外部唯一键单号
func (r *AlibabaServicecenterWorkcardCreateRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetOuterId() string {
    return r._outerId
}
// ServiceProvider Setter
// 服务提供者信息
func (r *AlibabaServicecenterWorkcardCreateRequest) SetServiceProvider(_serviceProvider *ServiceProviderDto) error {
    r._serviceProvider = _serviceProvider
    r.Set("service_provider", _serviceProvider)
    return nil
}

// ServiceProvider Getter
func (r AlibabaServicecenterWorkcardCreateRequest) GetServiceProvider() *ServiceProviderDto {
    return r._serviceProvider
}
