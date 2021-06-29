package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取饿了么用户信息 API请求
taobao.miniapp.eleuserinfo.get

获取饿了么用户信息
*/
type TaobaoMiniappEleuserinfoGetRequest struct {
    model.Params
    // 怪兽业务方
    bizScence   string
}

// 初始化TaobaoMiniappEleuserinfoGetRequest对象
func NewTaobaoMiniappEleuserinfoGetRequest() *TaobaoMiniappEleuserinfoGetRequest{
    return &TaobaoMiniappEleuserinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappEleuserinfoGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.eleuserinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappEleuserinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizScence Setter
// 怪兽业务方
func (r *TaobaoMiniappEleuserinfoGetRequest) SetBizScence(bizScence string) error {
    r.bizScence = bizScence
    r.Set("biz_scence", bizScence)
    return nil
}

// BizScence Getter
func (r TaobaoMiniappEleuserinfoGetRequest) GetBizScence() string {
    return r.bizScence
}
