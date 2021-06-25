package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
供应商编辑产品API 
taobao.alitrip.travel.product.base.modify

飞猪供销平台供应商可通过该API编辑产品
*/
func TaobaoAlitripTravelProductBaseModify(clt *core.SDKClient, req *travel.TaobaoAlitripTravelProductBaseModifyRequest, session string) (*travel.TaobaoAlitripTravelProductBaseModifyResponse, error) {
    var resp travel.TaobaoAlitripTravelProductBaseModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
