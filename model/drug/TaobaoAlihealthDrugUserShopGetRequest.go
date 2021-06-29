package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户id获取店铺id API请求
taobao.alihealth.drug.user.shop.get

提供给千牛智能客服，获取用户当前咨询的店铺ID
*/
type TaobaoAlihealthDrugUserShopGetRequest struct {
    model.Params
    // 用户昵称
    userNick   string
}

// 初始化TaobaoAlihealthDrugUserShopGetRequest对象
func NewTaobaoAlihealthDrugUserShopGetRequest() *TaobaoAlihealthDrugUserShopGetRequest{
    return &TaobaoAlihealthDrugUserShopGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugUserShopGetRequest) GetApiMethodName() string {
    return "taobao.alihealth.drug.user.shop.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugUserShopGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 用户昵称
func (r *TaobaoAlihealthDrugUserShopGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r TaobaoAlihealthDrugUserShopGetRequest) GetUserNick() string {
    return r.userNick
}
