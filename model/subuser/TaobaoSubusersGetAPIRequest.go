package subuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersGetAPIRequest 获取指定账户的子账号简易信息列表 API请求
// taobao.subusers.get
//
// 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
type TaobaoSubusersGetAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
}

// NewTaobaoSubusersGetRequest 初始化TaobaoSubusersGetAPIRequest对象
func NewTaobaoSubusersGetRequest() *TaobaoSubusersGetAPIRequest {
	return &TaobaoSubusersGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubusersGetAPIRequest) Reset() {
	r._userNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubusersGetAPIRequest) GetApiMethodName() string {
	return "taobao.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubusersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubusersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 主账号用户名
func (r *TaobaoSubusersGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoSubusersGetAPIRequest) GetUserNick() string {
	return r._userNick
}

var poolTaobaoSubusersGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubusersGetRequest()
	},
}

// GetTaobaoSubusersGetRequest 从 sync.Pool 获取 TaobaoSubusersGetAPIRequest
func GetTaobaoSubusersGetAPIRequest() *TaobaoSubusersGetAPIRequest {
	return poolTaobaoSubusersGetAPIRequest.Get().(*TaobaoSubusersGetAPIRequest)
}

// ReleaseTaobaoSubusersGetAPIRequest 将 TaobaoSubusersGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubusersGetAPIRequest(v *TaobaoSubusersGetAPIRequest) {
	v.Reset()
	poolTaobaoSubusersGetAPIRequest.Put(v)
}
