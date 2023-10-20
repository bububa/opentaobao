package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytcodereplaceAPIRequest 单码替换 API请求
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
type AlibabaalihealthdrugkytcodereplaceAPIRequest struct {
	model.Params
	// 企业ref_ent_id（码申请企业）
	_refEntId string
	// 替换后的追溯码
	_newCode string
	// 被替换的追溯码
	_oldCode string
}

// NewAlibabaalihealthdrugkytcodereplaceRequest 初始化AlibabaalihealthdrugkytcodereplaceAPIRequest对象
func NewAlibabaalihealthdrugkytcodereplaceRequest() *AlibabaalihealthdrugkytcodereplaceAPIRequest {
	return &AlibabaalihealthdrugkytcodereplaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytcodereplaceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codereplace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytcodereplaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytcodereplaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ref_ent_id（码申请企业）
func (r *AlibabaalihealthdrugkytcodereplaceAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytcodereplaceAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetNewCode is NewCode Setter
// 替换后的追溯码
func (r *AlibabaalihealthdrugkytcodereplaceAPIRequest) SetNewCode(_newCode string) error {
	r._newCode = _newCode
	r.Set("new_code", _newCode)
	return nil
}

// GetNewCode NewCode Getter
func (r AlibabaalihealthdrugkytcodereplaceAPIRequest) GetNewCode() string {
	return r._newCode
}

// SetOldCode is OldCode Setter
// 被替换的追溯码
func (r *AlibabaalihealthdrugkytcodereplaceAPIRequest) SetOldCode(_oldCode string) error {
	r._oldCode = _oldCode
	r.Set("old_code", _oldCode)
	return nil
}

// GetOldCode OldCode Getter
func (r AlibabaalihealthdrugkytcodereplaceAPIRequest) GetOldCode() string {
	return r._oldCode
}
