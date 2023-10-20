package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercentersubusersgetAPIRequest 查询指定账户的子账号列表 API请求
// taobao.sellercenter.subusers.get
//
// 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
type TaobaosellercentersubusersgetAPIRequest struct {
	model.Params
	// 表示卖家昵称
	_nick string
}

// NewTaobaosellercentersubusersgetRequest 初始化TaobaosellercentersubusersgetAPIRequest对象
func NewTaobaosellercentersubusersgetRequest() *TaobaosellercentersubusersgetAPIRequest {
	return &TaobaosellercentersubusersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosellercentersubusersgetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosellercentersubusersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosellercentersubusersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 表示卖家昵称
func (r *TaobaosellercentersubusersgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosellercentersubusersgetAPIRequest) GetNick() string {
	return r._nick
}
