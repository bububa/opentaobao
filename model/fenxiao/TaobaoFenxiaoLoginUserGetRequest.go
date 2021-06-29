package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分销用户登录信息 API请求
taobao.fenxiao.login.user.get

获取用户登录信息
*/
type TaobaoFenxiaoLoginUserGetRequest struct {
    model.Params
}

// 初始化TaobaoFenxiaoLoginUserGetRequest对象
func NewTaobaoFenxiaoLoginUserGetRequest() *TaobaoFenxiaoLoginUserGetRequest{
    return &TaobaoFenxiaoLoginUserGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoLoginUserGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.login.user.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoLoginUserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
