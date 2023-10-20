package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrsinglerelationAPIRequest 多融单码关联关系查询 API请求
// alibaba.alihealth.drug.kyt.dr.singlerelation
//
// 单码关联关系查询
type AlibabaalihealthdrugkytdrsinglerelationAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 追溯码
	_code string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaalihealthdrugkytdrsinglerelationRequest 初始化AlibabaalihealthdrugkytdrsinglerelationAPIRequest对象
func NewAlibabaalihealthdrugkytdrsinglerelationRequest() *AlibabaalihealthdrugkytdrsinglerelationAPIRequest {
	return &AlibabaalihealthdrugkytdrsinglerelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrsinglerelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.singlerelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrsinglerelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrsinglerelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaalihealthdrugkytdrsinglerelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrsinglerelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugkytdrsinglerelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytdrsinglerelationAPIRequest) GetCode() string {
	return r._code
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaalihealthdrugkytdrsinglerelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaalihealthdrugkytdrsinglerelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}
