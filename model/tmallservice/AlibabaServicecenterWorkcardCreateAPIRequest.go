package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardcreateAPIRequest 服务平台工单创建接口 API请求
// alibaba.servicecenter.workcard.create
//
// 创建服务平台工单
type AlibabaservicecenterworkcardcreateAPIRequest struct {
	model.Params
	// 工单属性，json格式字符串
	_attributes string
	// 工单外部唯一键单号
	_outerId string
	// 服务单id
	_spServiceOrderId int64
	// 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
	_serviceSequence int64
	// 申请次数
	_serviceCount int64
	// 服务提供者信息
	_serviceProvider *ServiceProviderDto
}

// NewAlibabaservicecenterworkcardcreateRequest 初始化AlibabaservicecenterworkcardcreateAPIRequest对象
func NewAlibabaservicecenterworkcardcreateRequest() *AlibabaservicecenterworkcardcreateAPIRequest {
	return &AlibabaservicecenterworkcardcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributes is Attributes Setter
// 工单属性，json格式字符串
func (r *AlibabaservicecenterworkcardcreateAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetOuterId is OuterId Setter
// 工单外部唯一键单号
func (r *AlibabaservicecenterworkcardcreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetSpServiceOrderId is SpServiceOrderId Setter
// 服务单id
func (r *AlibabaservicecenterworkcardcreateAPIRequest) SetSpServiceOrderId(_spServiceOrderId int64) error {
	r._spServiceOrderId = _spServiceOrderId
	r.Set("sp_service_order_id", _spServiceOrderId)
	return nil
}

// GetSpServiceOrderId SpServiceOrderId Getter
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetSpServiceOrderId() int64 {
	return r._spServiceOrderId
}

// SetServiceSequence is ServiceSequence Setter
// 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
func (r *AlibabaservicecenterworkcardcreateAPIRequest) SetServiceSequence(_serviceSequence int64) error {
	r._serviceSequence = _serviceSequence
	r.Set("service_sequence", _serviceSequence)
	return nil
}

// GetServiceSequence ServiceSequence Getter
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetServiceSequence() int64 {
	return r._serviceSequence
}

// SetServiceCount is ServiceCount Setter
// 申请次数
func (r *AlibabaservicecenterworkcardcreateAPIRequest) SetServiceCount(_serviceCount int64) error {
	r._serviceCount = _serviceCount
	r.Set("service_count", _serviceCount)
	return nil
}

// GetServiceCount ServiceCount Getter
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetServiceCount() int64 {
	return r._serviceCount
}

// SetServiceProvider is ServiceProvider Setter
// 服务提供者信息
func (r *AlibabaservicecenterworkcardcreateAPIRequest) SetServiceProvider(_serviceProvider *ServiceProviderDto) error {
	r._serviceProvider = _serviceProvider
	r.Set("service_provider", _serviceProvider)
	return nil
}

// GetServiceProvider ServiceProvider Getter
func (r AlibabaservicecenterworkcardcreateAPIRequest) GetServiceProvider() *ServiceProviderDto {
	return r._serviceProvider
}
