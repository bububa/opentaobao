package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubusersinfoqueryAPIRequest 根据当前子账号登陆态，获取该子账号基本信息 API请求
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
type TaobaosubusersinfoqueryAPIRequest struct {
	model.Params
	// 站点信息，淘宝天猫传0，1688传3
	_site int64
}

// NewTaobaosubusersinfoqueryRequest 初始化TaobaosubusersinfoqueryAPIRequest对象
func NewTaobaosubusersinfoqueryRequest() *TaobaosubusersinfoqueryAPIRequest {
	return &TaobaosubusersinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubusersinfoqueryAPIRequest) GetApiMethodName() string {
	return "taobao.subusers.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubusersinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubusersinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSite is Site Setter
// 站点信息，淘宝天猫传0，1688传3
func (r *TaobaosubusersinfoqueryAPIRequest) SetSite(_site int64) error {
	r._site = _site
	r.Set("site", _site)
	return nil
}

// GetSite Site Getter
func (r TaobaosubusersinfoqueryAPIRequest) GetSite() int64 {
	return r._site
}
