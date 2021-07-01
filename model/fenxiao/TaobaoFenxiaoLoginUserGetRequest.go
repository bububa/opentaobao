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
type TaobaoFenxiaoLoginUserGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoFenxiaoLoginUserGetAPIRequest对象
func NewTaobaoFenxiaoLoginUserGetRequest() *TaobaoFenxiaoLoginUserGetAPIRequest{
    return &TaobaoFenxiaoLoginUserGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoLoginUserGetAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.login.user.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoLoginUserGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
