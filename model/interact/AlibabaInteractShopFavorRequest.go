package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺收藏 API请求
alibaba.interact.shop.favor

店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
*/
type AlibabaInteractShopFavorRequest struct {
    model.Params
}

// 初始化AlibabaInteractShopFavorRequest对象
func NewAlibabaInteractShopFavorRequest() *AlibabaInteractShopFavorRequest{
    return &AlibabaInteractShopFavorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractShopFavorRequest) GetApiMethodName() string {
    return "alibaba.interact.shop.favor"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractShopFavorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
