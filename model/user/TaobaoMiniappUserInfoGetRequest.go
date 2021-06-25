package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户开放信息获取 APIRequest
taobao.miniapp.userInfo.get

获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
*/
type TaobaoMiniappUserInfoGetRequest struct {
    model.Params

}

func NewTaobaoMiniappUserInfoGetRequest() *TaobaoMiniappUserInfoGetRequest{
    return &TaobaoMiniappUserInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappUserInfoGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.userInfo.get"
}

func (r TaobaoMiniappUserInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


