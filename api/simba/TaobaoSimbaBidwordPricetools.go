package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
关键词出价指导工具（新） 
taobao.simba.bidword.pricetools

关键词出价指导工具（新）
*/
func TaobaoSimbaBidwordPricetools(clt *core.SDKClient, req *simba.TaobaoSimbaBidwordPricetoolsRequest, session string) (*simba.TaobaoSimbaBidwordPricetoolsAPIResponse, error) {
    var resp simba.TaobaoSimbaBidwordPricetoolsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
