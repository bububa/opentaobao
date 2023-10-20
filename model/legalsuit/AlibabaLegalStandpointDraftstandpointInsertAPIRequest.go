package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointDraftstandpointInsertAPIRequest 编辑后新增草稿口径 API请求
// alibaba.legal.standpoint.draftstandpoint.insert
//
// 编辑后新增草稿口径
type AlibabaLegalStandpointDraftstandpointInsertAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 口径
	_standpointDeriveDraftEpaasDto *StandpointDeriveDraftEpaasDto
}

// NewAlibabaLegalStandpointDraftstandpointInsertRequest 初始化AlibabaLegalStandpointDraftstandpointInsertAPIRequest对象
func NewAlibabaLegalStandpointDraftstandpointInsertRequest() *AlibabaLegalStandpointDraftstandpointInsertAPIRequest {
	return &AlibabaLegalStandpointDraftstandpointInsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointDraftstandpointInsertAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.draftstandpoint.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointDraftstandpointInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointDraftstandpointInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointDraftstandpointInsertAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointDraftstandpointInsertAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetStandpointDeriveDraftEpaasDto is StandpointDeriveDraftEpaasDto Setter
// 口径
func (r *AlibabaLegalStandpointDraftstandpointInsertAPIRequest) SetStandpointDeriveDraftEpaasDto(_standpointDeriveDraftEpaasDto *StandpointDeriveDraftEpaasDto) error {
	r._standpointDeriveDraftEpaasDto = _standpointDeriveDraftEpaasDto
	r.Set("standpoint_derive_draft_epaas_dto", _standpointDeriveDraftEpaasDto)
	return nil
}

// GetStandpointDeriveDraftEpaasDto StandpointDeriveDraftEpaasDto Getter
func (r AlibabaLegalStandpointDraftstandpointInsertAPIRequest) GetStandpointDeriveDraftEpaasDto() *StandpointDeriveDraftEpaasDto {
	return r._standpointDeriveDraftEpaasDto
}
