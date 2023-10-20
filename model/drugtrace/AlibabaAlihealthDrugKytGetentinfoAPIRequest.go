package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetentinfoAPIRequest 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API请求
// alibaba.alihealth.drug.kyt.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaalihealthdrugkytgetentinfoAPIRequest struct {
	model.Params
	// 企业名称
	_entName string
}

// NewAlibabaalihealthdrugkytgetentinfoRequest 初始化AlibabaalihealthdrugkytgetentinfoAPIRequest对象
func NewAlibabaalihealthdrugkytgetentinfoRequest() *AlibabaalihealthdrugkytgetentinfoAPIRequest {
	return &AlibabaalihealthdrugkytgetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytgetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytgetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytgetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaalihealthdrugkytgetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytgetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
