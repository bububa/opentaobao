package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaToolsItemsTopGetAPIRequest
取得一个关键词的推广组排名列表 API请求
taobao.simba.tools.items.top.get

取得一个关键词的推广组排名列表 */
type TaobaoSimbaToolsItemsTopGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词
	_keyword string
	// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
	_ip string
}

// NewTaobaoSimbaToolsItemsTopGetRequest 初始化TaobaoSimbaToolsItemsTopGetAPIRequest对象
func NewTaobaoSimbaToolsItemsTopGetRequest() *TaobaoSimbaToolsItemsTopGetAPIRequest {
	return &TaobaoSimbaToolsItemsTopGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.tools.items.top.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaToolsItemsTopGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is Keyword Setter
// 关键词
func (r *TaobaoSimbaToolsItemsTopGetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// Get Keyword Getter
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetKeyword() string {
	return r._keyword
}

// Set is Ip Setter
// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
func (r *TaobaoSimbaToolsItemsTopGetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// Get Ip Getter
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetIp() string {
	return r._ip
}
