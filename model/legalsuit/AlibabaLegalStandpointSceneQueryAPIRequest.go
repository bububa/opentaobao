package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointscenequeryAPIRequest 查询场景 API请求
// alibaba.legal.standpoint.scene.query
//
// 查询场景
type AlibabalegalstandpointscenequeryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
}

// NewAlibabalegalstandpointscenequeryRequest 初始化AlibabalegalstandpointscenequeryAPIRequest对象
func NewAlibabalegalstandpointscenequeryRequest() *AlibabalegalstandpointscenequeryAPIRequest {
	return &AlibabalegalstandpointscenequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointscenequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.scene.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointscenequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointscenequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointscenequeryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointscenequeryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}
