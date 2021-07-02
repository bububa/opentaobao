package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterSubusersGetAPIRequest 查询指定账户的子账号列表 API请求
// taobao.sellercenter.subusers.get
//
// 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
type TaobaoSellercenterSubusersGetAPIRequest struct {
	model.Params
	// 表示卖家昵称
	_nick string
}

// NewTaobaoSellercenterSubusersGetRequest 初始化TaobaoSellercenterSubusersGetAPIRequest对象
func NewTaobaoSellercenterSubusersGetRequest() *TaobaoSellercenterSubusersGetAPIRequest {
	return &TaobaoSellercenterSubusersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterSubusersGetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterSubusersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 表示卖家昵称
func (r *TaobaoSellercenterSubusersGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSellercenterSubusersGetAPIRequest) GetNick() string {
	return r._nick
}
