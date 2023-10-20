package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesgetentinfoAPIRequest 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API请求
// alibaba.alihealth.drug.kyt.wes.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaalihealthdrugkytwesgetentinfoAPIRequest struct {
	model.Params
	// 企业名称
	_entName string
}

// NewAlibabaalihealthdrugkytwesgetentinfoRequest 初始化AlibabaalihealthdrugkytwesgetentinfoAPIRequest对象
func NewAlibabaalihealthdrugkytwesgetentinfoRequest() *AlibabaalihealthdrugkytwesgetentinfoAPIRequest {
	return &AlibabaalihealthdrugkytwesgetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytwesgetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytwesgetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytwesgetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaalihealthdrugkytwesgetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytwesgetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
