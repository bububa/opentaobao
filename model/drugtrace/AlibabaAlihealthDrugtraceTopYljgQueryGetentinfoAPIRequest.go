package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest 通过企业名得到唯一标识ref_ent_id及企业ent_id API请求
// alibaba.alihealth.drugtrace.top.yljg.query.getentinfo
//
// 根据企业名称查询ID
type AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaalihealthdrugtracetopyljgquerygetentinfoRequest 初始化AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest对象
func NewAlibabaalihealthdrugtracetopyljgquerygetentinfoRequest() *AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest {
	return &AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
