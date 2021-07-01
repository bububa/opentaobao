package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户开放信息获取 API请求
taobao.miniapp.userInfo.get

获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
*/
type TaobaoMiniappUserInfoGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoMiniappUserInfoGetAPIRequest对象
func NewTaobaoMiniappUserInfoGetRequest() *TaobaoMiniappUserInfoGetAPIRequest{
    return &TaobaoMiniappUserInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappUserInfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.userInfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappUserInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
