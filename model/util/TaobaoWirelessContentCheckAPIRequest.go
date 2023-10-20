package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowirelesscontentcheckAPIRequest 无线开放内容检查 API请求
// taobao.wireless.content.check
//
// 无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 &lt;a href=&#34;https://help.aliyun.com/document_detail/70439.html&#34; target=&#34;blank&#34;&gt;阿里云内容安全&lt;/a&gt;
type TaobaowirelesscontentcheckAPIRequest struct {
	model.Params
	// 待检查的文本
	_text string
}

// NewTaobaowirelesscontentcheckRequest 初始化TaobaowirelesscontentcheckAPIRequest对象
func NewTaobaowirelesscontentcheckRequest() *TaobaowirelesscontentcheckAPIRequest {
	return &TaobaowirelesscontentcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowirelesscontentcheckAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.content.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowirelesscontentcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowirelesscontentcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetText is Text Setter
// 待检查的文本
func (r *TaobaowirelesscontentcheckAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TaobaowirelesscontentcheckAPIRequest) GetText() string {
	return r._text
}
