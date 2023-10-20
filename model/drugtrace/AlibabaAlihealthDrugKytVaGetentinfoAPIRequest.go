package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytvagetentinfoAPIRequest 通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) API请求
// alibaba.alihealth.drug.kyt.va.getentinfo
//
// 根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
type AlibabaalihealthdrugkytvagetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaalihealthdrugkytvagetentinfoRequest 初始化AlibabaalihealthdrugkytvagetentinfoAPIRequest对象
func NewAlibabaalihealthdrugkytvagetentinfoRequest() *AlibabaalihealthdrugkytvagetentinfoAPIRequest {
	return &AlibabaalihealthdrugkytvagetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytvagetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.va.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytvagetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytvagetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaalihealthdrugkytvagetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytvagetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
