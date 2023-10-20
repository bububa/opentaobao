package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweitaofollowisrelationAPIRequest 微淘是否关注 API请求
// taobao.weitao.follow.isrelation
//
// 判断用户是否关注对应的公共账号
type TaobaoweitaofollowisrelationAPIRequest struct {
	model.Params
	// 要查询的粉丝的淘宝昵称
	_fansNick string
	// ouid
	_ouid string
	// openid
	_openid string
}

// NewTaobaoweitaofollowisrelationRequest 初始化TaobaoweitaofollowisrelationAPIRequest对象
func NewTaobaoweitaofollowisrelationRequest() *TaobaoweitaofollowisrelationAPIRequest {
	return &TaobaoweitaofollowisrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweitaofollowisrelationAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.follow.isrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweitaofollowisrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweitaofollowisrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFansNick is FansNick Setter
// 要查询的粉丝的淘宝昵称
func (r *TaobaoweitaofollowisrelationAPIRequest) SetFansNick(_fansNick string) error {
	r._fansNick = _fansNick
	r.Set("fans_nick", _fansNick)
	return nil
}

// GetFansNick FansNick Getter
func (r TaobaoweitaofollowisrelationAPIRequest) GetFansNick() string {
	return r._fansNick
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoweitaofollowisrelationAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoweitaofollowisrelationAPIRequest) GetOuid() string {
	return r._ouid
}

// SetOpenid is Openid Setter
// openid
func (r *TaobaoweitaofollowisrelationAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r TaobaoweitaofollowisrelationAPIRequest) GetOpenid() string {
	return r._openid
}
