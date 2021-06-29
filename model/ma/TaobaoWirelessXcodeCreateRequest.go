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
    bizCode   string
    // 原始内容/原始地址
    content   string
    // 普通二维码样式参数
    style   *QrCodeStyle
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
func (r *TaobaoWirelessXcodeCreateRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r TaobaoWirelessXcodeCreateRequest) GetBizCode() string {
    return r.bizCode
}
// Content Setter
// 原始内容/原始地址
func (r *TaobaoWirelessXcodeCreateRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r TaobaoWirelessXcodeCreateRequest) GetContent() string {
    return r.content
}
// Style Setter
// 普通二维码样式参数
func (r *TaobaoWirelessXcodeCreateRequest) SetStyle(style *QrCodeStyle) error {
    r.style = style
    r.Set("style", style)
    return nil
}

// Style Getter
func (r TaobaoWirelessXcodeCreateRequest) GetStyle() *QrCodeStyle {
    return r.style
}
