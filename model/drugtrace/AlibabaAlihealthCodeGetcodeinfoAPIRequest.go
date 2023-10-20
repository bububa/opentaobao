package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthcodegetcodeinfoAPIRequest 码查询功能 API请求
// alibaba.alihealth.code.getcodeinfo
//
// 码查询功能
type AlibabaalihealthcodegetcodeinfoAPIRequest struct {
	model.Params
	// 企业refEntId
	_refEntId string
	// 码
	_code string
}

// NewAlibabaalihealthcodegetcodeinfoRequest 初始化AlibabaalihealthcodegetcodeinfoAPIRequest对象
func NewAlibabaalihealthcodegetcodeinfoRequest() *AlibabaalihealthcodegetcodeinfoAPIRequest {
	return &AlibabaalihealthcodegetcodeinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthcodegetcodeinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.code.getcodeinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthcodegetcodeinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthcodegetcodeinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaalihealthcodegetcodeinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthcodegetcodeinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 码
func (r *AlibabaalihealthcodegetcodeinfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthcodegetcodeinfoAPIRequest) GetCode() string {
	return r._code
}
