package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessContentCheckAPIRequest 无线开放内容检查 API请求
// taobao.wireless.content.check
//
// 无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 &lt;a href=&#34;https://help.aliyun.com/document_detail/70439.html&#34; target=&#34;blank&#34;&gt;阿里云内容安全&lt;/a&gt;
type TaobaoWirelessContentCheckAPIRequest struct {
	model.Params
	// 待检查的文本
	_text string
}

// NewTaobaoWirelessContentCheckRequest 初始化TaobaoWirelessContentCheckAPIRequest对象
func NewTaobaoWirelessContentCheckRequest() *TaobaoWirelessContentCheckAPIRequest {
	return &TaobaoWirelessContentCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWirelessContentCheckAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.content.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWirelessContentCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetText is Text Setter
// 待检查的文本
func (r *TaobaoWirelessContentCheckAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TaobaoWirelessContentCheckAPIRequest) GetText() string {
	return r._text
}
