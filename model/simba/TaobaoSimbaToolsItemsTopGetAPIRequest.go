package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaToolsItemsTopGetAPIRequest 取得一个关键词的推广组排名列表 API请求
// taobao.simba.tools.items.top.get
//
// 取得一个关键词的推广组排名列表
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaToolsItemsTopGetAPIRequest) Reset() {
	r._nick = ""
	r._keyword = ""
	r._ip = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.tools.items.top.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaToolsItemsTopGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetNick() string {
	return r._nick
}

// SetKeyword is Keyword Setter
// 关键词
func (r *TaobaoSimbaToolsItemsTopGetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetIp is Ip Setter
// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
func (r *TaobaoSimbaToolsItemsTopGetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaoSimbaToolsItemsTopGetAPIRequest) GetIp() string {
	return r._ip
}

var poolTaobaoSimbaToolsItemsTopGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaToolsItemsTopGetRequest()
	},
}

// GetTaobaoSimbaToolsItemsTopGetRequest 从 sync.Pool 获取 TaobaoSimbaToolsItemsTopGetAPIRequest
func GetTaobaoSimbaToolsItemsTopGetAPIRequest() *TaobaoSimbaToolsItemsTopGetAPIRequest {
	return poolTaobaoSimbaToolsItemsTopGetAPIRequest.Get().(*TaobaoSimbaToolsItemsTopGetAPIRequest)
}

// ReleaseTaobaoSimbaToolsItemsTopGetAPIRequest 将 TaobaoSimbaToolsItemsTopGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaToolsItemsTopGetAPIRequest(v *TaobaoSimbaToolsItemsTopGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaToolsItemsTopGetAPIRequest.Put(v)
}
