package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
乐动力线-微信设备绑定通知 API请求
alibaba.alisports.ldlserver.wxevent.post

转发弹外微信服务到弹内
*/
type AlibabaAlisportsLdlserverWxeventPostAPIRequest struct {
    model.Params
    // 微信服务器消息数据
    _rawstr   string
}

// 初始化AlibabaAlisportsLdlserverWxeventPostAPIRequest对象
func NewAlibabaAlisportsLdlserverWxeventPostRequest() *AlibabaAlisportsLdlserverWxeventPostAPIRequest{
    return &AlibabaAlisportsLdlserverWxeventPostAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsLdlserverWxeventPostAPIRequest) GetApiMethodName() string {
    return "alibaba.alisports.ldlserver.wxevent.post"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsLdlserverWxeventPostAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rawstr Setter
// 微信服务器消息数据
func (r *AlibabaAlisportsLdlserverWxeventPostAPIRequest) SetRawstr(_rawstr string) error {
    r._rawstr = _rawstr
    r.Set("rawstr", _rawstr)
    return nil
}

// Rawstr Getter
func (r AlibabaAlisportsLdlserverWxeventPostAPIRequest) GetRawstr() string {
    return r._rawstr
}
