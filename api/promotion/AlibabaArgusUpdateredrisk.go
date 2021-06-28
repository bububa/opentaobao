package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
更新红线价格 
alibaba.argus.updateredrisk

商品健康中心新增红线价格规则
*/
func AlibabaArgusUpdateredrisk(clt *core.SDKClient, req *promotion.AlibabaArgusUpdateredriskRequest, session string) (*promotion.AlibabaArgusUpdateredriskAPIResponse, error) {
    var resp promotion.AlibabaArgusUpdateredriskAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
