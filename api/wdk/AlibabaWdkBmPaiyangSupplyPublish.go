package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
派样商品库存变更同步接口 
alibaba.wdk.bm.paiyang.supply.publish

淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
*/
func AlibabaWdkBmPaiyangSupplyPublish(clt *core.SDKClient, req *wdk.AlibabaWdkBmPaiyangSupplyPublishAPIRequest, session string) (*wdk.AlibabaWdkBmPaiyangSupplyPublishAPIResponse, error) {
    var resp wdk.AlibabaWdkBmPaiyangSupplyPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
