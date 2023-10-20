package subuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubuserDutysGetAPIRequest 获取指定账户的所有职务信息列表 API请求
// taobao.subuser.dutys.get
//
// 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
type TaobaoSubuserDutysGetAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
}

// NewTaobaoSubuserDutysGetRequest 初始化TaobaoSubuserDutysGetAPIRequest对象
func NewTaobaoSubuserDutysGetRequest() *TaobaoSubuserDutysGetAPIRequest {
	return &TaobaoSubuserDutysGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubuserDutysGetAPIRequest) Reset() {
	r._userNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubuserDutysGetAPIRequest) GetApiMethodName() string {
	return "taobao.subuser.dutys.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubuserDutysGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubuserDutysGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 主账号用户名
func (r *TaobaoSubuserDutysGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoSubuserDutysGetAPIRequest) GetUserNick() string {
	return r._userNick
}

var poolTaobaoSubuserDutysGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubuserDutysGetRequest()
	},
}

// GetTaobaoSubuserDutysGetRequest 从 sync.Pool 获取 TaobaoSubuserDutysGetAPIRequest
func GetTaobaoSubuserDutysGetAPIRequest() *TaobaoSubuserDutysGetAPIRequest {
	return poolTaobaoSubuserDutysGetAPIRequest.Get().(*TaobaoSubuserDutysGetAPIRequest)
}

// ReleaseTaobaoSubuserDutysGetAPIRequest 将 TaobaoSubuserDutysGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubuserDutysGetAPIRequest(v *TaobaoSubuserDutysGetAPIRequest) {
	v.Reset()
	poolTaobaoSubuserDutysGetAPIRequest.Put(v)
}
