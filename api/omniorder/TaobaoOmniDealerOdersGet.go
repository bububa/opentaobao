package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
获取单笔全渠道经销商订单的详细信息 
taobao.omni.dealer.oders.get

全渠道经销商获取单笔订单的详细信息
*/
func TaobaoOmniDealerOdersGet(clt *core.SDKClient, req *omniorder.TaobaoOmniDealerOdersGetRequest, session string) (*omniorder.TaobaoOmniDealerOdersGetAPIResponse, error) {
    var resp omniorder.TaobaoOmniDealerOdersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
