package subuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterUserPermissionsGetAPIRequest 获取指定用户的权限集合 API请求
// taobao.sellercenter.user.permissions.get
//
// 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
type TaobaoSellercenterUserPermissionsGetAPIRequest struct {
	model.Params
	// 用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。
	_nick string
}

// NewTaobaoSellercenterUserPermissionsGetRequest 初始化TaobaoSellercenterUserPermissionsGetAPIRequest对象
func NewTaobaoSellercenterUserPermissionsGetRequest() *TaobaoSellercenterUserPermissionsGetAPIRequest {
	return &TaobaoSellercenterUserPermissionsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSellercenterUserPermissionsGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterUserPermissionsGetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.user.permissions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterUserPermissionsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSellercenterUserPermissionsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。
func (r *TaobaoSellercenterUserPermissionsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSellercenterUserPermissionsGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoSellercenterUserPermissionsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSellercenterUserPermissionsGetRequest()
	},
}

// GetTaobaoSellercenterUserPermissionsGetRequest 从 sync.Pool 获取 TaobaoSellercenterUserPermissionsGetAPIRequest
func GetTaobaoSellercenterUserPermissionsGetAPIRequest() *TaobaoSellercenterUserPermissionsGetAPIRequest {
	return poolTaobaoSellercenterUserPermissionsGetAPIRequest.Get().(*TaobaoSellercenterUserPermissionsGetAPIRequest)
}

// ReleaseTaobaoSellercenterUserPermissionsGetAPIRequest 将 TaobaoSellercenterUserPermissionsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSellercenterUserPermissionsGetAPIRequest(v *TaobaoSellercenterUserPermissionsGetAPIRequest) {
	v.Reset()
	poolTaobaoSellercenterUserPermissionsGetAPIRequest.Put(v)
}
