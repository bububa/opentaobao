package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺收藏 APIRequest
alibaba.interact.shop.favor

店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
*/
type AlibabaInteractShopFavorRequest struct {
    model.Params

}

func NewAlibabaInteractShopFavorRequest() *AlibabaInteractShopFavorRequest{
    return &AlibabaInteractShopFavorRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractShopFavorRequest) GetApiMethodName() string {
    return "alibaba.interact.shop.favor"
}

func (r AlibabaInteractShopFavorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


