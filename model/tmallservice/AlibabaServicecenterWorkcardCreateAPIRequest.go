package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardCreateAPIRequest 服务平台工单创建接口 API请求
// alibaba.servicecenter.workcard.create
//
// 创建服务平台工单
type AlibabaServicecenterWorkcardCreateAPIRequest struct {
	model.Params
	// 服务单id
	_spServiceOrderId int64
	// 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
	_serviceSequence int64
	// 申请次数
	_serviceCount int64
	// 工单属性，json格式字符串
	_attributes string
	// 工单外部唯一键单号
	_outerId string
	// 服务提供者信息
	_serviceProvider *ServiceProviderDto
}

// NewAlibabaServicecenterWorkcardCreateRequest 初始化AlibabaServicecenterWorkcardCreateAPIRequest对象
func NewAlibabaServicecenterWorkcardCreateRequest() *AlibabaServicecenterWorkcardCreateAPIRequest {
	return &AlibabaServicecenterWorkcardCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSpServiceOrderId is SpServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterWorkcardCreateAPIRequest) SetSpServiceOrderId(_spServiceOrderId int64) error {
	r._spServiceOrderId = _spServiceOrderId
	r.Set("sp_service_order_id", _spServiceOrderId)
	return nil
}

// GetSpServiceOrderId SpServiceOrderId Getter
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetSpServiceOrderId() int64 {
	return r._spServiceOrderId
}

// SetServiceSequence is ServiceSequence Setter
// 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
func (r *AlibabaServicecenterWorkcardCreateAPIRequest) SetServiceSequence(_serviceSequence int64) error {
	r._serviceSequence = _serviceSequence
	r.Set("service_sequence", _serviceSequence)
	return nil
}

// GetServiceSequence ServiceSequence Getter
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetServiceSequence() int64 {
	return r._serviceSequence
}

// SetServiceCount is ServiceCount Setter
// 申请次数
func (r *AlibabaServicecenterWorkcardCreateAPIRequest) SetServiceCount(_serviceCount int64) error {
	r._serviceCount = _serviceCount
	r.Set("service_count", _serviceCount)
	return nil
}

// GetServiceCount ServiceCount Getter
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetServiceCount() int64 {
	return r._serviceCount
}

// SetAttributes is Attributes Setter
// 工单属性，json格式字符串
func (r *AlibabaServicecenterWorkcardCreateAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetOuterId is OuterId Setter
// 工单外部唯一键单号
func (r *AlibabaServicecenterWorkcardCreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetServiceProvider is ServiceProvider Setter
// 服务提供者信息
func (r *AlibabaServicecenterWorkcardCreateAPIRequest) SetServiceProvider(_serviceProvider *ServiceProviderDto) error {
	r._serviceProvider = _serviceProvider
	r.Set("service_provider", _serviceProvider)
	return nil
}

// GetServiceProvider ServiceProvider Getter
func (r AlibabaServicecenterWorkcardCreateAPIRequest) GetServiceProvider() *ServiceProviderDto {
	return r._serviceProvider
}
