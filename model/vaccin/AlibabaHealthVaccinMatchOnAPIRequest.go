package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinMatchOnAPIRequest isv自主上下线疫苗，可以选择上线还是下线 API请求
// alibaba.health.vaccin.match.on
//
// isv自主上下线疫苗，可以选择上线还是下线
type AlibabaHealthVaccinMatchOnAPIRequest struct {
	model.Params
	// 阿里健康的povId
	_povId string
	// 阿里健康的vaccineId
	_vaccineId string
	// 对码类型：ON表示上线，OFF表示下线
	_isvMatchType string
}

// NewAlibabaHealthVaccinMatchOnRequest 初始化AlibabaHealthVaccinMatchOnAPIRequest对象
func NewAlibabaHealthVaccinMatchOnRequest() *AlibabaHealthVaccinMatchOnAPIRequest {
	return &AlibabaHealthVaccinMatchOnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinMatchOnAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.match.on"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinMatchOnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthVaccinMatchOnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPovId is PovId Setter
// 阿里健康的povId
func (r *AlibabaHealthVaccinMatchOnAPIRequest) SetPovId(_povId string) error {
	r._povId = _povId
	r.Set("pov_id", _povId)
	return nil
}

// GetPovId PovId Getter
func (r AlibabaHealthVaccinMatchOnAPIRequest) GetPovId() string {
	return r._povId
}

// SetVaccineId is VaccineId Setter
// 阿里健康的vaccineId
func (r *AlibabaHealthVaccinMatchOnAPIRequest) SetVaccineId(_vaccineId string) error {
	r._vaccineId = _vaccineId
	r.Set("vaccine_id", _vaccineId)
	return nil
}

// GetVaccineId VaccineId Getter
func (r AlibabaHealthVaccinMatchOnAPIRequest) GetVaccineId() string {
	return r._vaccineId
}

// SetIsvMatchType is IsvMatchType Setter
// 对码类型：ON表示上线，OFF表示下线
func (r *AlibabaHealthVaccinMatchOnAPIRequest) SetIsvMatchType(_isvMatchType string) error {
	r._isvMatchType = _isvMatchType
	r.Set("isv_match_type", _isvMatchType)
	return nil
}

// GetIsvMatchType IsvMatchType Getter
func (r AlibabaHealthVaccinMatchOnAPIRequest) GetIsvMatchType() string {
	return r._isvMatchType
}
