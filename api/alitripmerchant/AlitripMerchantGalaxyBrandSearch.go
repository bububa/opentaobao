package alitripmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripmerchant"
)

/* 
星河-品牌搜索 
alitrip.merchant.galaxy.brand.search

星河服务=获取雅高品牌信息
*/
func AlitripMerchantGalaxyBrandSearch(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyBrandSearchAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyBrandSearchAPIResponse, error) {
    var resp alitripmerchant.AlitripMerchantGalaxyBrandSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
