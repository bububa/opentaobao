package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
限时打折详情查询 
taobao.promotion.limitdiscount.detail.get

限时打折详情查询。查询出指定限时打折的对应商品记录信息。
*/
func TaobaoPromotionLimitdiscountDetailGet(clt *core.SDKClient, req *promotion.TaobaoPromotionLimitdiscountDetailGetRequest, session string) (*promotion.TaobaoPromotionLimitdiscountDetailGetResponse, error) {
    var resp promotion.TaobaoPromotionLimitdiscountDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
