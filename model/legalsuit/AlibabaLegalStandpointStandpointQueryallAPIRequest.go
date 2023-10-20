package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointstandpointqueryallAPIRequest 滑动查询口径 API请求
// alibaba.legal.standpoint.standpoint.queryall
//
// 滑动查询口径
type AlibabalegalstandpointstandpointqueryallAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 滑动查询参数
	_queryParam *QueryParam
}

// NewAlibabalegalstandpointstandpointqueryallRequest 初始化AlibabalegalstandpointstandpointqueryallAPIRequest对象
func NewAlibabalegalstandpointstandpointqueryallRequest() *AlibabalegalstandpointstandpointqueryallAPIRequest {
	return &AlibabalegalstandpointstandpointqueryallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointstandpointqueryallAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpoint.queryall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointstandpointqueryallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointstandpointqueryallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointstandpointqueryallAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointstandpointqueryallAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetQueryParam is QueryParam Setter
// 滑动查询参数
func (r *AlibabalegalstandpointstandpointqueryallAPIRequest) SetQueryParam(_queryParam *QueryParam) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r AlibabalegalstandpointstandpointqueryallAPIRequest) GetQueryParam() *QueryParam {
	return r._queryParam
}
