package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetentlicenseAPIRequest 获取企业资质 API请求
// alibaba.alihealth.drug.kyt.getentlicense
//
// 获取企业的资质信息。
type AlibabaalihealthdrugkytgetentlicenseAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
}

// NewAlibabaalihealthdrugkytgetentlicenseRequest 初始化AlibabaalihealthdrugkytgetentlicenseAPIRequest对象
func NewAlibabaalihealthdrugkytgetentlicenseRequest() *AlibabaalihealthdrugkytgetentlicenseAPIRequest {
	return &AlibabaalihealthdrugkytgetentlicenseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytgetentlicenseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getentlicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytgetentlicenseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytgetentlicenseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytgetentlicenseAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytgetentlicenseAPIRequest) GetRefEntId() string {
	return r._refEntId
}
