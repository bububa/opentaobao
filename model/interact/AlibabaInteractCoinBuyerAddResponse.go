package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
平台向买家发放淘金币 APIResponse
alibaba.interact.coin.buyer.add

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。ISV调用该接口向买家发放平台淘金币，需要优惠平台运营审核ISV资质。
*/
type AlibabaInteractCoinBuyerAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractCoinBuyerAddResponse `json:"alibaba_interact_coin_buyer_add_response,omitempty"`
}

type AlibabaInteractCoinBuyerAddResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
