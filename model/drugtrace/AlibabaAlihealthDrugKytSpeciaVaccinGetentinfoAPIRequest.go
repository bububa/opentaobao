package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest 通过企业名得到唯一标识(ref_ent_id)及企业ID(ent_id) API请求
// alibaba.alihealth.drug.kyt.specia.vaccin.getentinfo
//
// 根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
type AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaalihealthdrugkytspeciavaccingetentinfoRequest 初始化AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest对象
func NewAlibabaalihealthdrugkytspeciavaccingetentinfoRequest() *AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest {
	return &AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.specia.vaccin.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytspeciavaccingetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
