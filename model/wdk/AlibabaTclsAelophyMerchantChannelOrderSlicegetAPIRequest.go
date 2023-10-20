package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelorderslicegetAPIRequest 获取运力时间片信息 API请求
// alibaba.tcls.aelophy.merchant.channel.order.sliceget
//
// 获取履约时间片
type AlibabatclsaelophymerchantchannelorderslicegetAPIRequest struct {
	model.Params
	// 获取时间片入参
	_timeSliceGetRequest *TimeSliceGetRequest
}

// NewAlibabatclsaelophymerchantchannelorderslicegetRequest 初始化AlibabatclsaelophymerchantchannelorderslicegetAPIRequest对象
func NewAlibabatclsaelophymerchantchannelorderslicegetRequest() *AlibabatclsaelophymerchantchannelorderslicegetAPIRequest {
	return &AlibabatclsaelophymerchantchannelorderslicegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelorderslicegetAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.sliceget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelorderslicegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelorderslicegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTimeSliceGetRequest is TimeSliceGetRequest Setter
// 获取时间片入参
func (r *AlibabatclsaelophymerchantchannelorderslicegetAPIRequest) SetTimeSliceGetRequest(_timeSliceGetRequest *TimeSliceGetRequest) error {
	r._timeSliceGetRequest = _timeSliceGetRequest
	r.Set("time_slice_get_request", _timeSliceGetRequest)
	return nil
}

// GetTimeSliceGetRequest TimeSliceGetRequest Getter
func (r AlibabatclsaelophymerchantchannelorderslicegetAPIRequest) GetTimeSliceGetRequest() *TimeSliceGetRequest {
	return r._timeSliceGetRequest
}
