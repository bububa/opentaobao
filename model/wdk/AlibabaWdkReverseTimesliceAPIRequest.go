package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseTimesliceAPIRequest 逆向取货时间片查询 API请求
// alibaba.wdk.reverse.timeslice
//
// 逆向取货时间片查询
type AlibabaWdkReverseTimesliceAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramQueryTimeSliceReq *QueryTimeSliceReq
}

// NewAlibabaWdkReverseTimesliceRequest 初始化AlibabaWdkReverseTimesliceAPIRequest对象
func NewAlibabaWdkReverseTimesliceRequest() *AlibabaWdkReverseTimesliceAPIRequest {
	return &AlibabaWdkReverseTimesliceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkReverseTimesliceAPIRequest) Reset() {
	r._paramQueryTimeSliceReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseTimesliceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.timeslice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseTimesliceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkReverseTimesliceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryTimeSliceReq is ParamQueryTimeSliceReq Setter
// 系统自动生成
func (r *AlibabaWdkReverseTimesliceAPIRequest) SetParamQueryTimeSliceReq(_paramQueryTimeSliceReq *QueryTimeSliceReq) error {
	r._paramQueryTimeSliceReq = _paramQueryTimeSliceReq
	r.Set("param_query_time_slice_req", _paramQueryTimeSliceReq)
	return nil
}

// GetParamQueryTimeSliceReq ParamQueryTimeSliceReq Getter
func (r AlibabaWdkReverseTimesliceAPIRequest) GetParamQueryTimeSliceReq() *QueryTimeSliceReq {
	return r._paramQueryTimeSliceReq
}

var poolAlibabaWdkReverseTimesliceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkReverseTimesliceRequest()
	},
}

// GetAlibabaWdkReverseTimesliceRequest 从 sync.Pool 获取 AlibabaWdkReverseTimesliceAPIRequest
func GetAlibabaWdkReverseTimesliceAPIRequest() *AlibabaWdkReverseTimesliceAPIRequest {
	return poolAlibabaWdkReverseTimesliceAPIRequest.Get().(*AlibabaWdkReverseTimesliceAPIRequest)
}

// ReleaseAlibabaWdkReverseTimesliceAPIRequest 将 AlibabaWdkReverseTimesliceAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkReverseTimesliceAPIRequest(v *AlibabaWdkReverseTimesliceAPIRequest) {
	v.Reset()
	poolAlibabaWdkReverseTimesliceAPIRequest.Put(v)
}
