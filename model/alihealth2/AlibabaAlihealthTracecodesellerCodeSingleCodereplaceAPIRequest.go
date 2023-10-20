package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest 非药单码替换 API请求
// alibaba.alihealth.tracecodeseller.code.single.codereplace
//
// 提供非药追溯码单码替换功能
type AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest struct {
	model.Params
	// 企业id
	_entInfoId string
	// 新码
	_newCode string
	// 老码
	_oldCode string
}

// NewAlibabaalihealthtracecodesellercodesinglecodereplaceRequest 初始化AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest对象
func NewAlibabaalihealthtracecodesellercodesinglecodereplaceRequest() *AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest {
	return &AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.code.single.codereplace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntInfoId is EntInfoId Setter
// 企业id
func (r *AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) SetEntInfoId(_entInfoId string) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) GetEntInfoId() string {
	return r._entInfoId
}

// SetNewCode is NewCode Setter
// 新码
func (r *AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) SetNewCode(_newCode string) error {
	r._newCode = _newCode
	r.Set("new_code", _newCode)
	return nil
}

// GetNewCode NewCode Getter
func (r AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) GetNewCode() string {
	return r._newCode
}

// SetOldCode is OldCode Setter
// 老码
func (r *AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) SetOldCode(_oldCode string) error {
	r._oldCode = _oldCode
	r.Set("old_code", _oldCode)
	return nil
}

// GetOldCode OldCode Getter
func (r AlibabaalihealthtracecodesellercodesinglecodereplaceAPIRequest) GetOldCode() string {
	return r._oldCode
}
