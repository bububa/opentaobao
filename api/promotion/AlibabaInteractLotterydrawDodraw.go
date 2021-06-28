package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
无线端抽奖接口 
alibaba.interact.lotterydraw.dodraw

商家抽奖平台无线端抽奖接口开放
*/
func AlibabaInteractLotterydrawDodraw(clt *core.SDKClient, req *promotion.AlibabaInteractLotterydrawDodrawRequest, session string) (*promotion.AlibabaInteractLotterydrawDodrawAPIResponse, error) {
    var resp promotion.AlibabaInteractLotterydrawDodrawAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
