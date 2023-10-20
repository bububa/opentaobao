package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqyputpackageAPIRequest 码拼箱 API请求
// alibaba.alihealth.drug.kyt.scqy.putpackage
//
// 码拼箱接口
type AlibabaalihealthdrugkytscqyputpackageAPIRequest struct {
	model.Params
	// 企业id
	_refEntid string
	// 二级码
	_secondaryCode string
	// 一级码
	_primaryCodes string
}

// NewAlibabaalihealthdrugkytscqyputpackageRequest 初始化AlibabaalihealthdrugkytscqyputpackageAPIRequest对象
func NewAlibabaalihealthdrugkytscqyputpackageRequest() *AlibabaalihealthdrugkytscqyputpackageAPIRequest {
	return &AlibabaalihealthdrugkytscqyputpackageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytscqyputpackageAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.putpackage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytscqyputpackageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytscqyputpackageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntid is RefEntid Setter
// 企业id
func (r *AlibabaalihealthdrugkytscqyputpackageAPIRequest) SetRefEntid(_refEntid string) error {
	r._refEntid = _refEntid
	r.Set("ref_entid", _refEntid)
	return nil
}

// GetRefEntid RefEntid Getter
func (r AlibabaalihealthdrugkytscqyputpackageAPIRequest) GetRefEntid() string {
	return r._refEntid
}

// SetSecondaryCode is SecondaryCode Setter
// 二级码
func (r *AlibabaalihealthdrugkytscqyputpackageAPIRequest) SetSecondaryCode(_secondaryCode string) error {
	r._secondaryCode = _secondaryCode
	r.Set("secondary_code", _secondaryCode)
	return nil
}

// GetSecondaryCode SecondaryCode Getter
func (r AlibabaalihealthdrugkytscqyputpackageAPIRequest) GetSecondaryCode() string {
	return r._secondaryCode
}

// SetPrimaryCodes is PrimaryCodes Setter
// 一级码
func (r *AlibabaalihealthdrugkytscqyputpackageAPIRequest) SetPrimaryCodes(_primaryCodes string) error {
	r._primaryCodes = _primaryCodes
	r.Set("primary_codes", _primaryCodes)
	return nil
}

// GetPrimaryCodes PrimaryCodes Getter
func (r AlibabaalihealthdrugkytscqyputpackageAPIRequest) GetPrimaryCodes() string {
	return r._primaryCodes
}
