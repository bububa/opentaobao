package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbatoolsitemstopgetAPIRequest 取得一个关键词的推广组排名列表 API请求
// taobao.simba.tools.items.top.get
//
// 取得一个关键词的推广组排名列表
type TaobaosimbatoolsitemstopgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词
	_keyword string
	// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
	_ip string
}

// NewTaobaosimbatoolsitemstopgetRequest 初始化TaobaosimbatoolsitemstopgetAPIRequest对象
func NewTaobaosimbatoolsitemstopgetRequest() *TaobaosimbatoolsitemstopgetAPIRequest {
	return &TaobaosimbatoolsitemstopgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbatoolsitemstopgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.tools.items.top.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbatoolsitemstopgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbatoolsitemstopgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbatoolsitemstopgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbatoolsitemstopgetAPIRequest) GetNick() string {
	return r._nick
}

// SetKeyword is Keyword Setter
// 关键词
func (r *TaobaosimbatoolsitemstopgetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r TaobaosimbatoolsitemstopgetAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetIp is Ip Setter
// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
func (r *TaobaosimbatoolsitemstopgetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaosimbatoolsitemstopgetAPIRequest) GetIp() string {
	return r._ip
}
