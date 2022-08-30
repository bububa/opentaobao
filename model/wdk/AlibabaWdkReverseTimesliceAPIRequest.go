package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseTimesliceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.timeslice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseTimesliceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
