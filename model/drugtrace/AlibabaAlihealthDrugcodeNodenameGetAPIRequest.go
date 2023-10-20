package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodenodenamegetAPIRequest 根据码获取机构名称 API请求
// alibaba.alihealth.drugcode.nodename.get
//
// 根据码获取机构名称
type AlibabaalihealthdrugcodenodenamegetAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业id
	_refEntId string
}

// NewAlibabaalihealthdrugcodenodenamegetRequest 初始化AlibabaalihealthdrugcodenodenamegetAPIRequest对象
func NewAlibabaalihealthdrugcodenodenamegetRequest() *AlibabaalihealthdrugcodenodenamegetAPIRequest {
	return &AlibabaalihealthdrugcodenodenamegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodenodenamegetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.nodename.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodenodenamegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodenodenamegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugcodenodenamegetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugcodenodenamegetAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugcodenodenamegetAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodenodenamegetAPIRequest) GetRefEntId() string {
	return r._refEntId
}
