package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫奖池鉴权接口 APIRequest
alibaba.interact.isvlottery.isvdraw

鉴权接口，为tida.isvdraw接口鉴权
*/
type AlibabaInteractIsvlotteryIsvdrawRequest struct {
    model.Params

}

func NewAlibabaInteractIsvlotteryIsvdrawRequest() *AlibabaInteractIsvlotteryIsvdrawRequest{
    return &AlibabaInteractIsvlotteryIsvdrawRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractIsvlotteryIsvdrawRequest) GetApiMethodName() string {
    return "alibaba.interact.isvlottery.isvdraw"
}

func (r AlibabaInteractIsvlotteryIsvdrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


