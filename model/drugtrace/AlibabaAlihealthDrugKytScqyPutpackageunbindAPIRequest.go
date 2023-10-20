package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest 码拼箱解除父子关系接口 API请求
// alibaba.alihealth.drug.kyt.scqy.putpackageunbind
//
// 码拼箱解除父子关系接口
type AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 解绑的父码
	_parentCode string
	// 解绑的子码
	_childCodes string
}

// NewAlibabaalihealthdrugkytscqyputpackageunbindRequest 初始化AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest对象
func NewAlibabaalihealthdrugkytscqyputpackageunbindRequest() *AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest {
	return &AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.putpackageunbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetParentCode is ParentCode Setter
// 解绑的父码
func (r *AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) SetParentCode(_parentCode string) error {
	r._parentCode = _parentCode
	r.Set("parent_code", _parentCode)
	return nil
}

// GetParentCode ParentCode Getter
func (r AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) GetParentCode() string {
	return r._parentCode
}

// SetChildCodes is ChildCodes Setter
// 解绑的子码
func (r *AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) SetChildCodes(_childCodes string) error {
	r._childCodes = _childCodes
	r.Set("child_codes", _childCodes)
	return nil
}

// GetChildCodes ChildCodes Getter
func (r AlibabaalihealthdrugkytscqyputpackageunbindAPIRequest) GetChildCodes() string {
	return r._childCodes
}
