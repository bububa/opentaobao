package ma

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessXcodeCreateAPIRequest 创建二维码/短连接 API请求
// taobao.wireless.xcode.create
//
// 创建码平台的普通二维码或者长连接转短连接服务
type TaobaoWirelessXcodeCreateAPIRequest struct {
	model.Params
	// 码平台开放的业务code
	_bizCode string
	// 原始内容/原始地址
	_content string
	// 普通二维码样式参数
	_style *QrCodeStyle
}

// NewTaobaoWirelessXcodeCreateRequest 初始化TaobaoWirelessXcodeCreateAPIRequest对象
func NewTaobaoWirelessXcodeCreateRequest() *TaobaoWirelessXcodeCreateAPIRequest {
	return &TaobaoWirelessXcodeCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWirelessXcodeCreateAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.xcode.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWirelessXcodeCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizCode is BizCode Setter
// 码平台开放的业务code
func (r *TaobaoWirelessXcodeCreateAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TaobaoWirelessXcodeCreateAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetContent is Content Setter
// 原始内容/原始地址
func (r *TaobaoWirelessXcodeCreateAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoWirelessXcodeCreateAPIRequest) GetContent() string {
	return r._content
}

// SetStyle is Style Setter
// 普通二维码样式参数
func (r *TaobaoWirelessXcodeCreateAPIRequest) SetStyle(_style *QrCodeStyle) error {
	r._style = _style
	r.Set("style", _style)
	return nil
}

// GetStyle Style Getter
func (r TaobaoWirelessXcodeCreateAPIRequest) GetStyle() *QrCodeStyle {
	return r._style
}
