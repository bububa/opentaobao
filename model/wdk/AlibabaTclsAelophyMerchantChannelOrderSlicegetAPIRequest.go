package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest 获取运力时间片信息 API请求
// alibaba.tcls.aelophy.merchant.channel.order.sliceget
//
// 获取履约时间片
type AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest struct {
	model.Params
	// 获取时间片入参
	_timeSliceGetRequest *TimeSliceGetRequest
}

// NewAlibabaTclsAelophyMerchantChannelOrderSlicegetRequest 初始化AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderSlicegetRequest() *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest) Reset() {
	r._timeSliceGetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.sliceget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTimeSliceGetRequest is TimeSliceGetRequest Setter
// 获取时间片入参
func (r *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest) SetTimeSliceGetRequest(_timeSliceGetRequest *TimeSliceGetRequest) error {
	r._timeSliceGetRequest = _timeSliceGetRequest
	r.Set("time_slice_get_request", _timeSliceGetRequest)
	return nil
}

// GetTimeSliceGetRequest TimeSliceGetRequest Getter
func (r AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest) GetTimeSliceGetRequest() *TimeSliceGetRequest {
	return r._timeSliceGetRequest
}

var poolAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantChannelOrderSlicegetRequest()
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderSlicegetRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest
func GetAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest() *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest {
	return poolAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest.Get().(*AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest 将 AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest(v *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest.Put(v)
}
