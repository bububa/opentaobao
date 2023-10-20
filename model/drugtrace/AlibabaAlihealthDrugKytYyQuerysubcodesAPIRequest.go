package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytyyquerysubcodesAPIRequest 查询一个码的所有子码 API请求
// alibaba.alihealth.drug.kyt.yy.querysubcodes
//
// 单码的了码查询
type AlibabaalihealthdrugkytyyquerysubcodesAPIRequest struct {
	model.Params
	// 码
	_codes []string
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
}

// NewAlibabaalihealthdrugkytyyquerysubcodesRequest 初始化AlibabaalihealthdrugkytyyquerysubcodesAPIRequest对象
func NewAlibabaalihealthdrugkytyyquerysubcodesRequest() *AlibabaalihealthdrugkytyyquerysubcodesAPIRequest {
	return &AlibabaalihealthdrugkytyyquerysubcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytyyquerysubcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.querysubcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytyyquerysubcodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytyyquerysubcodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码
func (r *AlibabaalihealthdrugkytyyquerysubcodesAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugkytyyquerysubcodesAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaalihealthdrugkytyyquerysubcodesAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytyyquerysubcodesAPIRequest) GetRefEntId() string {
	return r._refEntId
}
