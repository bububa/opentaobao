package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqyputpackagebindAPIRequest 码拼箱建立父子关系接口 API请求
// alibaba.alihealth.drug.kyt.scqy.putpackagebind
//
// 码拼箱建立父子关系接口
type AlibabaalihealthdrugkytscqyputpackagebindAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 替换父码
	_sourceParentCode string
	// 替换子码
	_sourceChildCodes string
	// 被替换父码
	_targetParentCode string
}

// NewAlibabaalihealthdrugkytscqyputpackagebindRequest 初始化AlibabaalihealthdrugkytscqyputpackagebindAPIRequest对象
func NewAlibabaalihealthdrugkytscqyputpackagebindRequest() *AlibabaalihealthdrugkytscqyputpackagebindAPIRequest {
	return &AlibabaalihealthdrugkytscqyputpackagebindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.putpackagebind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetSourceParentCode is SourceParentCode Setter
// 替换父码
func (r *AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) SetSourceParentCode(_sourceParentCode string) error {
	r._sourceParentCode = _sourceParentCode
	r.Set("source_parent_code", _sourceParentCode)
	return nil
}

// GetSourceParentCode SourceParentCode Getter
func (r AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) GetSourceParentCode() string {
	return r._sourceParentCode
}

// SetSourceChildCodes is SourceChildCodes Setter
// 替换子码
func (r *AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) SetSourceChildCodes(_sourceChildCodes string) error {
	r._sourceChildCodes = _sourceChildCodes
	r.Set("source_child_codes", _sourceChildCodes)
	return nil
}

// GetSourceChildCodes SourceChildCodes Getter
func (r AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) GetSourceChildCodes() string {
	return r._sourceChildCodes
}

// SetTargetParentCode is TargetParentCode Setter
// 被替换父码
func (r *AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) SetTargetParentCode(_targetParentCode string) error {
	r._targetParentCode = _targetParentCode
	r.Set("target_parent_code", _targetParentCode)
	return nil
}

// GetTargetParentCode TargetParentCode Getter
func (r AlibabaalihealthdrugkytscqyputpackagebindAPIRequest) GetTargetParentCode() string {
	return r._targetParentCode
}
