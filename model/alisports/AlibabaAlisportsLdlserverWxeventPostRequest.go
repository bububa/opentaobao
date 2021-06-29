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
type AlibabaAlisportsLdlserverWxeventPostRequest struct {
    model.Params
    // 微信服务器消息数据
    rawstr   string
}

// 初始化AlibabaAlisportsLdlserverWxeventPostRequest对象
func NewAlibabaAlisportsLdlserverWxeventPostRequest() *AlibabaAlisportsLdlserverWxeventPostRequest{
    return &AlibabaAlisportsLdlserverWxeventPostRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsLdlserverWxeventPostRequest) GetApiMethodName() string {
    return "alibaba.alisports.ldlserver.wxevent.post"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsLdlserverWxeventPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rawstr Setter
// 微信服务器消息数据
func (r *AlibabaAlisportsLdlserverWxeventPostRequest) SetRawstr(rawstr string) error {
    r.rawstr = rawstr
    r.Set("rawstr", rawstr)
    return nil
}

// Rawstr Getter
func (r AlibabaAlisportsLdlserverWxeventPostRequest) GetRawstr() string {
    return r.rawstr
}
