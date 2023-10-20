package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointdraftstandpointinsertAPIRequest 编辑后新增草稿口径 API请求
// alibaba.legal.standpoint.draftstandpoint.insert
//
// 编辑后新增草稿口径
type AlibabalegalstandpointdraftstandpointinsertAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 口径
	_standpointDeriveDraftEpaasDto *StandpointDeriveDraftEpaasDto
}

// NewAlibabalegalstandpointdraftstandpointinsertRequest 初始化AlibabalegalstandpointdraftstandpointinsertAPIRequest对象
func NewAlibabalegalstandpointdraftstandpointinsertRequest() *AlibabalegalstandpointdraftstandpointinsertAPIRequest {
	return &AlibabalegalstandpointdraftstandpointinsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointdraftstandpointinsertAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.draftstandpoint.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointdraftstandpointinsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointdraftstandpointinsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointdraftstandpointinsertAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointdraftstandpointinsertAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetStandpointDeriveDraftEpaasDto is StandpointDeriveDraftEpaasDto Setter
// 口径
func (r *AlibabalegalstandpointdraftstandpointinsertAPIRequest) SetStandpointDeriveDraftEpaasDto(_standpointDeriveDraftEpaasDto *StandpointDeriveDraftEpaasDto) error {
	r._standpointDeriveDraftEpaasDto = _standpointDeriveDraftEpaasDto
	r.Set("standpoint_derive_draft_epaas_dto", _standpointDeriveDraftEpaasDto)
	return nil
}

// GetStandpointDeriveDraftEpaasDto StandpointDeriveDraftEpaasDto Getter
func (r AlibabalegalstandpointdraftstandpointinsertAPIRequest) GetStandpointDeriveDraftEpaasDto() *StandpointDeriveDraftEpaasDto {
	return r._standpointDeriveDraftEpaasDto
}
