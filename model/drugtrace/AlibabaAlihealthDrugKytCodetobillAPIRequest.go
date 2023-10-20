package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytcodetobillAPIRequest 通过追溯码查单据 API请求
// alibaba.alihealth.drug.kyt.codetobill
//
// 通过追溯码查单据
type AlibabaalihealthdrugkytcodetobillAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 追溯码
	_code string
}

// NewAlibabaalihealthdrugkytcodetobillRequest 初始化AlibabaalihealthdrugkytcodetobillAPIRequest对象
func NewAlibabaalihealthdrugkytcodetobillRequest() *AlibabaalihealthdrugkytcodetobillAPIRequest {
	return &AlibabaalihealthdrugkytcodetobillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytcodetobillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codetobill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytcodetobillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytcodetobillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytcodetobillAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytcodetobillAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugkytcodetobillAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytcodetobillAPIRequest) GetCode() string {
	return r._code
}
