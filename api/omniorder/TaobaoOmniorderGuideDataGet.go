package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
获取全渠道导购产品数据 
taobao.omniorder.guide.data.get

获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
*/
func TaobaoOmniorderGuideDataGet(clt *core.SDKClient, req *omniorder.TaobaoOmniorderGuideDataGetRequest, session string) (*omniorder.TaobaoOmniorderGuideDataGetAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderGuideDataGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
