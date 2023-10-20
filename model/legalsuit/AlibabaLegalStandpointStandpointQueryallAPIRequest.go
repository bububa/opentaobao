package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointQueryallAPIRequest 滑动查询口径 API请求
// alibaba.legal.standpoint.standpoint.queryall
//
// 滑动查询口径
type AlibabaLegalStandpointStandpointQueryallAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 滑动查询参数
	_queryParam *QueryParam
}

// NewAlibabaLegalStandpointStandpointQueryallRequest 初始化AlibabaLegalStandpointStandpointQueryallAPIRequest对象
func NewAlibabaLegalStandpointStandpointQueryallRequest() *AlibabaLegalStandpointStandpointQueryallAPIRequest {
	return &AlibabaLegalStandpointStandpointQueryallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointStandpointQueryallAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpoint.queryall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointStandpointQueryallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointStandpointQueryallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointStandpointQueryallAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointStandpointQueryallAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetQueryParam is QueryParam Setter
// 滑动查询参数
func (r *AlibabaLegalStandpointStandpointQueryallAPIRequest) SetQueryParam(_queryParam *QueryParam) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r AlibabaLegalStandpointStandpointQueryallAPIRequest) GetQueryParam() *QueryParam {
	return r._queryParam
}
