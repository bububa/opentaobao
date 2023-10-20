package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointstandpointqueryAPIRequest 查询具体口径 API请求
// alibaba.legal.standpoint.standpoint.query
//
// 查询具体口径
type AlibabalegalstandpointstandpointqueryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 口径id
	_standpointId int64
}

// NewAlibabalegalstandpointstandpointqueryRequest 初始化AlibabalegalstandpointstandpointqueryAPIRequest对象
func NewAlibabalegalstandpointstandpointqueryRequest() *AlibabalegalstandpointstandpointqueryAPIRequest {
	return &AlibabalegalstandpointstandpointqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointstandpointqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointstandpointqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointstandpointqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointstandpointqueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointstandpointqueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetStandpointId is StandpointId Setter
// 口径id
func (r *AlibabalegalstandpointstandpointqueryAPIRequest) SetStandpointId(_standpointId int64) error {
	r._standpointId = _standpointId
	r.Set("standpoint_id", _standpointId)
	return nil
}

// GetStandpointId StandpointId Getter
func (r AlibabalegalstandpointstandpointqueryAPIRequest) GetStandpointId() int64 {
	return r._standpointId
}
