package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytyygetentinfoAPIRequest 得到企业信息 API请求
// alibaba.alihealth.drug.kyt.yy.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaalihealthdrugkytyygetentinfoAPIRequest struct {
	model.Params
	// 公司名称
	_entName string
}

// NewAlibabaalihealthdrugkytyygetentinfoRequest 初始化AlibabaalihealthdrugkytyygetentinfoAPIRequest对象
func NewAlibabaalihealthdrugkytyygetentinfoRequest() *AlibabaalihealthdrugkytyygetentinfoAPIRequest {
	return &AlibabaalihealthdrugkytyygetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytyygetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytyygetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytyygetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称
func (r *AlibabaalihealthdrugkytyygetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytyygetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
