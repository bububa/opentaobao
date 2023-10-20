package ma

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowirelessxcodecreateAPIRequest 创建二维码/短连接 API请求
// taobao.wireless.xcode.create
//
// 创建码平台的普通二维码或者长连接转短连接服务
type TaobaowirelessxcodecreateAPIRequest struct {
	model.Params
	// 码平台开放的业务code
	_bizCode string
	// 原始内容/原始地址
	_content string
	// 普通二维码样式参数
	_style *QrCodeStyle
}

// NewTaobaowirelessxcodecreateRequest 初始化TaobaowirelessxcodecreateAPIRequest对象
func NewTaobaowirelessxcodecreateRequest() *TaobaowirelessxcodecreateAPIRequest {
	return &TaobaowirelessxcodecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowirelessxcodecreateAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.xcode.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowirelessxcodecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowirelessxcodecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 码平台开放的业务code
func (r *TaobaowirelessxcodecreateAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TaobaowirelessxcodecreateAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetContent is Content Setter
// 原始内容/原始地址
func (r *TaobaowirelessxcodecreateAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaowirelessxcodecreateAPIRequest) GetContent() string {
	return r._content
}

// SetStyle is Style Setter
// 普通二维码样式参数
func (r *TaobaowirelessxcodecreateAPIRequest) SetStyle(_style *QrCodeStyle) error {
	r._style = _style
	r.Set("style", _style)
	return nil
}

// GetStyle Style Getter
func (r TaobaowirelessxcodecreateAPIRequest) GetStyle() *QrCodeStyle {
	return r._style
}
