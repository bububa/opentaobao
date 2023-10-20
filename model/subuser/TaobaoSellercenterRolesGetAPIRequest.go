package subuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterRolesGetAPIRequest 获取指定卖家的角色列表 API请求
// taobao.sellercenter.roles.get
//
// 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
type TaobaoSellercenterRolesGetAPIRequest struct {
	model.Params
	// 卖家昵称(只允许查询自己的信息：当前登陆者)
	_nick string
}

// NewTaobaoSellercenterRolesGetRequest 初始化TaobaoSellercenterRolesGetAPIRequest对象
func NewTaobaoSellercenterRolesGetRequest() *TaobaoSellercenterRolesGetAPIRequest {
	return &TaobaoSellercenterRolesGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSellercenterRolesGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterRolesGetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.roles.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterRolesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSellercenterRolesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 卖家昵称(只允许查询自己的信息：当前登陆者)
func (r *TaobaoSellercenterRolesGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSellercenterRolesGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoSellercenterRolesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSellercenterRolesGetRequest()
	},
}

// GetTaobaoSellercenterRolesGetRequest 从 sync.Pool 获取 TaobaoSellercenterRolesGetAPIRequest
func GetTaobaoSellercenterRolesGetAPIRequest() *TaobaoSellercenterRolesGetAPIRequest {
	return poolTaobaoSellercenterRolesGetAPIRequest.Get().(*TaobaoSellercenterRolesGetAPIRequest)
}

// ReleaseTaobaoSellercenterRolesGetAPIRequest 将 TaobaoSellercenterRolesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSellercenterRolesGetAPIRequest(v *TaobaoSellercenterRolesGetAPIRequest) {
	v.Reset()
	poolTaobaoSellercenterRolesGetAPIRequest.Put(v)
}
