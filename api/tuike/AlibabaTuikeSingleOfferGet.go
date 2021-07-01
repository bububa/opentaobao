package tuike

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tuike"
)

/* 
查询1688推客平台卖家推广中的商品信息 
alibaba.tuike.single.offer.get

查询单个推客商品信息的接口
*/
func AlibabaTuikeSingleOfferGet(clt *core.SDKClient, req *tuike.AlibabaTuikeSingleOfferGetAPIRequest, session string) (*tuike.AlibabaTuikeSingleOfferGetAPIResponse, error) {
    var resp tuike.AlibabaTuikeSingleOfferGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
