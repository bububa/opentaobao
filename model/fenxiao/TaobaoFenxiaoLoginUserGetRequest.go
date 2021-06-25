package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分销用户登录信息 APIRequest
taobao.fenxiao.login.user.get

获取用户登录信息
*/
type TaobaoFenxiaoLoginUserGetRequest struct {
    model.Params

}

func NewTaobaoFenxiaoLoginUserGetRequest() *TaobaoFenxiaoLoginUserGetRequest{
    return &TaobaoFenxiaoLoginUserGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoLoginUserGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.login.user.get"
}

func (r TaobaoFenxiaoLoginUserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


