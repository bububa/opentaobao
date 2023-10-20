package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFollowIsrelationAPIRequest 微淘是否关注 API请求
// taobao.weitao.follow.isrelation
//
// 判断用户是否关注对应的公共账号
type TaobaoWeitaoFollowIsrelationAPIRequest struct {
	model.Params
	// 要查询的粉丝的淘宝昵称
	_fansNick string
	// ouid
	_ouid string
	// openid
	_openid string
}

// NewTaobaoWeitaoFollowIsrelationRequest 初始化TaobaoWeitaoFollowIsrelationAPIRequest对象
func NewTaobaoWeitaoFollowIsrelationRequest() *TaobaoWeitaoFollowIsrelationAPIRequest {
	return &TaobaoWeitaoFollowIsrelationAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeitaoFollowIsrelationAPIRequest) Reset() {
	r._fansNick = ""
	r._ouid = ""
	r._openid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFollowIsrelationAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.follow.isrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFollowIsrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeitaoFollowIsrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFansNick is FansNick Setter
// 要查询的粉丝的淘宝昵称
func (r *TaobaoWeitaoFollowIsrelationAPIRequest) SetFansNick(_fansNick string) error {
	r._fansNick = _fansNick
	r.Set("fans_nick", _fansNick)
	return nil
}

// GetFansNick FansNick Getter
func (r TaobaoWeitaoFollowIsrelationAPIRequest) GetFansNick() string {
	return r._fansNick
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoWeitaoFollowIsrelationAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoWeitaoFollowIsrelationAPIRequest) GetOuid() string {
	return r._ouid
}

// SetOpenid is Openid Setter
// openid
func (r *TaobaoWeitaoFollowIsrelationAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r TaobaoWeitaoFollowIsrelationAPIRequest) GetOpenid() string {
	return r._openid
}

var poolTaobaoWeitaoFollowIsrelationAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeitaoFollowIsrelationRequest()
	},
}

// GetTaobaoWeitaoFollowIsrelationRequest 从 sync.Pool 获取 TaobaoWeitaoFollowIsrelationAPIRequest
func GetTaobaoWeitaoFollowIsrelationAPIRequest() *TaobaoWeitaoFollowIsrelationAPIRequest {
	return poolTaobaoWeitaoFollowIsrelationAPIRequest.Get().(*TaobaoWeitaoFollowIsrelationAPIRequest)
}

// ReleaseTaobaoWeitaoFollowIsrelationAPIRequest 将 TaobaoWeitaoFollowIsrelationAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeitaoFollowIsrelationAPIRequest(v *TaobaoWeitaoFollowIsrelationAPIRequest) {
	v.Reset()
	poolTaobaoWeitaoFollowIsrelationAPIRequest.Put(v)
}
