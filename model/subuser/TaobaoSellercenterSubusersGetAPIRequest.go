package subuser

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSellercenterSubusersGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterSubusersGetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterSubusersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSellercenterSubusersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoSellercenterSubusersGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSellercenterSubusersGetRequest()
	},
}

// GetTaobaoSellercenterSubusersGetRequest 从 sync.Pool 获取 TaobaoSellercenterSubusersGetAPIRequest
func GetTaobaoSellercenterSubusersGetAPIRequest() *TaobaoSellercenterSubusersGetAPIRequest {
	return poolTaobaoSellercenterSubusersGetAPIRequest.Get().(*TaobaoSellercenterSubusersGetAPIRequest)
}

// ReleaseTaobaoSellercenterSubusersGetAPIRequest 将 TaobaoSellercenterSubusersGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSellercenterSubusersGetAPIRequest(v *TaobaoSellercenterSubusersGetAPIRequest) {
	v.Reset()
	poolTaobaoSellercenterSubusersGetAPIRequest.Put(v)
}
