package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceinventoryqueryAPIRequest 服务库存查询 API请求
// alibaba.ssc.supplyplatform.service.inventory.query
//
// 查询服务库存。需要保存服务容量成功后，才能进行查询，参数中的provider信息（provider_id和provider_type）与alibaba.ssc.supplyplatform.servicecapacity.save接口中保持一致。
type AlibabasscsupplyplatformserviceinventoryqueryAPIRequest struct {
	model.Params
	// 查询开始日期。yyyy-MM-dd格式
	_startDay string
	// 查询结束日期。与start间隔不能超过31天。yyyy-MM-dd格式
	_endDay string
	// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
	_providerType string
	// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
	_providerId int64
}

// NewAlibabasscsupplyplatformserviceinventoryqueryRequest 初始化AlibabasscsupplyplatformserviceinventoryqueryAPIRequest对象
func NewAlibabasscsupplyplatformserviceinventoryqueryRequest() *AlibabasscsupplyplatformserviceinventoryqueryAPIRequest {
	return &AlibabasscsupplyplatformserviceinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.service.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDay is StartDay Setter
// 查询开始日期。yyyy-MM-dd格式
func (r *AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) SetStartDay(_startDay string) error {
	r._startDay = _startDay
	r.Set("start_day", _startDay)
	return nil
}

// GetStartDay StartDay Getter
func (r AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) GetStartDay() string {
	return r._startDay
}

// SetEndDay is EndDay Setter
// 查询结束日期。与start间隔不能超过31天。yyyy-MM-dd格式
func (r *AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) SetEndDay(_endDay string) error {
	r._endDay = _endDay
	r.Set("end_day", _endDay)
	return nil
}

// GetEndDay EndDay Getter
func (r AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) GetEndDay() string {
	return r._endDay
}

// SetProviderType is ProviderType Setter
// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// GetProviderType ProviderType Getter
func (r AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) GetProviderType() string {
	return r._providerType
}

// SetProviderId is ProviderId Setter
// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabasscsupplyplatformserviceinventoryqueryAPIRequest) GetProviderId() int64 {
	return r._providerId
}
