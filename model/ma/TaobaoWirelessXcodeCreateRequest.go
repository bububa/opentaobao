package ma

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建二维码/短连接 API请求
taobao.wireless.xcode.create

创建码平台的普通二维码或者长连接转短连接服务
*/
type TaobaoWirelessXcodeCreateRequest struct {
    model.Params
    // 码平台开放的业务code
    _bizCode   string
    // 原始内容/原始地址
    _content   string
    // 普通二维码样式参数
    _style   *QrCodeStyle
}

// 初始化TaobaoWirelessXcodeCreateRequest对象
func NewTaobaoWirelessXcodeCreateRequest() *TaobaoWirelessXcodeCreateRequest{
    return &TaobaoWirelessXcodeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWirelessXcodeCreateRequest) GetApiMethodName() string {
    return "taobao.wireless.xcode.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWirelessXcodeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizCode Setter
// 码平台开放的业务code
func (r *TaobaoWirelessXcodeCreateRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r TaobaoWirelessXcodeCreateRequest) GetBizCode() string {
    return r._bizCode
}
// Content Setter
// 原始内容/原始地址
func (r *TaobaoWirelessXcodeCreateRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoWirelessXcodeCreateRequest) GetContent() string {
    return r._content
}
// Style Setter
// 普通二维码样式参数
func (r *TaobaoWirelessXcodeCreateRequest) SetStyle(_style *QrCodeStyle) error {
    r._style = _style
    r.Set("style", _style)
    return nil
}

// Style Getter
func (r TaobaoWirelessXcodeCreateRequest) GetStyle() *QrCodeStyle {
    return r._style
}
