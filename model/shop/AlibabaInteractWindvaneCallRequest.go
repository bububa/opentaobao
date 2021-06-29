package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Weex页面调用windvane API请求
alibaba.interact.windvane.call

客户端鉴权使用，实际不会发送或接收数据
*/
type AlibabaInteractWindvaneCallRequest struct {
    model.Params
    // 客户端鉴权使用，实际不会发送或接收数据
    _unNamed   string
}

// 初始化AlibabaInteractWindvaneCallRequest对象
func NewAlibabaInteractWindvaneCallRequest() *AlibabaInteractWindvaneCallRequest{
    return &AlibabaInteractWindvaneCallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractWindvaneCallRequest) GetApiMethodName() string {
    return "alibaba.interact.windvane.call"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractWindvaneCallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权使用，实际不会发送或接收数据
func (r *AlibabaInteractWindvaneCallRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaInteractWindvaneCallRequest) GetUnNamed() string {
    return r._unNamed
}
