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
type TaobaoAlihealthDrugUserShopGetAPIRequest struct {
    model.Params
    // 用户昵称
    _userNick   string
}

// 初始化TaobaoAlihealthDrugUserShopGetAPIRequest对象
func NewTaobaoAlihealthDrugUserShopGetRequest() *TaobaoAlihealthDrugUserShopGetAPIRequest{
    return &TaobaoAlihealthDrugUserShopGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugUserShopGetAPIRequest) GetApiMethodName() string {
    return "taobao.alihealth.drug.user.shop.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugUserShopGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 用户昵称
func (r *TaobaoAlihealthDrugUserShopGetAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoAlihealthDrugUserShopGetAPIRequest) GetUserNick() string {
    return r._userNick
}
