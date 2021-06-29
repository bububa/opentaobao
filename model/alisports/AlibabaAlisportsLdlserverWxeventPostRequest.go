package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
乐动力线-微信设备绑定通知 APIRequest
alibaba.alisports.ldlserver.wxevent.post

转发弹外微信服务到弹内
*/
type AlibabaAlisportsLdlserverWxeventPostRequest struct {
    model.Params

    // 微信服务器消息数据
    rawstr   string 

}

func NewAlibabaAlisportsLdlserverWxeventPostRequest() *AlibabaAlisportsLdlserverWxeventPostRequest{
    return &AlibabaAlisportsLdlserverWxeventPostRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsLdlserverWxeventPostRequest) GetApiMethodName() string {
    return "alibaba.alisports.ldlserver.wxevent.post"
}

func (r AlibabaAlisportsLdlserverWxeventPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsLdlserverWxeventPostRequest) SetRawstr(rawstr string) error {
    r.rawstr = rawstr
    r.Set("rawstr", rawstr)
    return nil
}

func (r AlibabaAlisportsLdlserverWxeventPostRequest) GetRawstr() string {
    return r.rawstr
}

