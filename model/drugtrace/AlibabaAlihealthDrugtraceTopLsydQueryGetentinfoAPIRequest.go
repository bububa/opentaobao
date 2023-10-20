package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest 通过企业名得到唯一标识ref_ent_id及企业ent_id API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo
//
// 根据企业名称查询ID
type AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaalihealthdrugtracetoplsydquerygetentinfoRequest 初始化AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest对象
func NewAlibabaalihealthdrugtracetoplsydquerygetentinfoRequest() *AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest {
	return &AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
