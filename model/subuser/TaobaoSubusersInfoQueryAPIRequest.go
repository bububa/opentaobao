package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersInfoQueryAPIRequest 根据当前子账号登陆态，获取该子账号基本信息 API请求
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
type TaobaoSubusersInfoQueryAPIRequest struct {
	model.Params
	// 站点信息，淘宝天猫传0，1688传3
	_site int64
}

// NewTaobaoSubusersInfoQueryRequest 初始化TaobaoSubusersInfoQueryAPIRequest对象
func NewTaobaoSubusersInfoQueryRequest() *TaobaoSubusersInfoQueryAPIRequest {
	return &TaobaoSubusersInfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubusersInfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.subusers.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubusersInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubusersInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSite is Site Setter
// 站点信息，淘宝天猫传0，1688传3
func (r *TaobaoSubusersInfoQueryAPIRequest) SetSite(_site int64) error {
	r._site = _site
	r.Set("site", _site)
	return nil
}

// GetSite Site Getter
func (r TaobaoSubusersInfoQueryAPIRequest) GetSite() int64 {
	return r._site
}
