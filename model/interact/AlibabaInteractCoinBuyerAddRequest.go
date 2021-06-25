package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
平台向买家发放淘金币 APIRequest
alibaba.interact.coin.buyer.add

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。ISV调用该接口向买家发放平台淘金币，需要优惠平台运营审核ISV资质。
*/
type AlibabaInteractCoinBuyerAddRequest struct {
    model.Params

}

func NewAlibabaInteractCoinBuyerAddRequest() *AlibabaInteractCoinBuyerAddRequest{
    return &AlibabaInteractCoinBuyerAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractCoinBuyerAddRequest) GetApiMethodName() string {
    return "alibaba.interact.coin.buyer.add"
}

func (r AlibabaInteractCoinBuyerAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


