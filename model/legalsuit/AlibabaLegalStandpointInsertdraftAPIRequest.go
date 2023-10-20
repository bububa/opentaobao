package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointinsertdraftAPIRequest 插入草稿 API请求
// alibaba.legal.standpoint.insertdraft
//
// 插入草稿
type AlibabalegalstandpointinsertdraftAPIRequest struct {
	model.Params
	// busid
	_busid string
	// busName
	_busName string
	// 系统标识
	_inputSystemCode string
	// 口径对象
	_standpointDraftDto *StandpointDraftDto
}

// NewAlibabalegalstandpointinsertdraftRequest 初始化AlibabalegalstandpointinsertdraftAPIRequest对象
func NewAlibabalegalstandpointinsertdraftRequest() *AlibabalegalstandpointinsertdraftAPIRequest {
	return &AlibabalegalstandpointinsertdraftAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointinsertdraftAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.insertdraft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointinsertdraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointinsertdraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusid is Busid Setter
// busid
func (r *AlibabalegalstandpointinsertdraftAPIRequest) SetBusid(_busid string) error {
	r._busid = _busid
	r.Set("busid", _busid)
	return nil
}

// GetBusid Busid Getter
func (r AlibabalegalstandpointinsertdraftAPIRequest) GetBusid() string {
	return r._busid
}

// SetBusName is BusName Setter
// busName
func (r *AlibabalegalstandpointinsertdraftAPIRequest) SetBusName(_busName string) error {
	r._busName = _busName
	r.Set("bus_name", _busName)
	return nil
}

// GetBusName BusName Getter
func (r AlibabalegalstandpointinsertdraftAPIRequest) GetBusName() string {
	return r._busName
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointinsertdraftAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointinsertdraftAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetStandpointDraftDto is StandpointDraftDto Setter
// 口径对象
func (r *AlibabalegalstandpointinsertdraftAPIRequest) SetStandpointDraftDto(_standpointDraftDto *StandpointDraftDto) error {
	r._standpointDraftDto = _standpointDraftDto
	r.Set("standpoint_draft_dto", _standpointDraftDto)
	return nil
}

// GetStandpointDraftDto StandpointDraftDto Getter
func (r AlibabalegalstandpointinsertdraftAPIRequest) GetStandpointDraftDto() *StandpointDraftDto {
	return r._standpointDraftDto
}
