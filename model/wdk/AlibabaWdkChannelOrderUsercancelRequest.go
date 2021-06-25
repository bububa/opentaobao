package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售中取消 APIRequest
alibaba.wdk.channel.order.usercancel

用户发起售中取消
*/
type AlibabaWdkChannelOrderUsercancelRequest struct {
    model.Params

    // 取消信息
    userCancelInfo   *OrderUserCancelInfo 

}

func NewAlibabaWdkChannelOrderUsercancelRequest() *AlibabaWdkChannelOrderUsercancelRequest{
    return &AlibabaWdkChannelOrderUsercancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkChannelOrderUsercancelRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.usercancel"
}

func (r AlibabaWdkChannelOrderUsercancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkChannelOrderUsercancelRequest) SetUserCancelInfo(userCancelInfo *OrderUserCancelInfo) error {
    r.userCancelInfo = userCancelInfo
    r.Set("user_cancel_info", userCancelInfo)
    return nil
}

func (r AlibabaWdkChannelOrderUsercancelRequest) GetUserCancelInfo() *OrderUserCancelInfo {
    return r.userCancelInfo
}

