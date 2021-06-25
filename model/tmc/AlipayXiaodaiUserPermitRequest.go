package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里金融用户授权 APIRequest
alipay.xiaodai.user.permit

阿里金融为用户开通消息通道接口
*/
type AlipayXiaodaiUserPermitRequest struct {
    model.Params

    // 用户数字ID
    userId   int64 

}

func NewAlipayXiaodaiUserPermitRequest() *AlipayXiaodaiUserPermitRequest{
    return &AlipayXiaodaiUserPermitRequest{
        Params: model.NewParams(),
    }
}

func (r AlipayXiaodaiUserPermitRequest) GetApiMethodName() string {
    return "alipay.xiaodai.user.permit"
}

func (r AlipayXiaodaiUserPermitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlipayXiaodaiUserPermitRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlipayXiaodaiUserPermitRequest) GetUserId() int64 {
    return r.userId
}

