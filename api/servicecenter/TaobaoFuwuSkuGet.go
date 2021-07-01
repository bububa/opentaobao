package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
获取内购服务及SKU详情 
taobao.fuwu.sku.get

通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息
*/
func TaobaoFuwuSkuGet(clt *core.SDKClient, req *servicecenter.TaobaoFuwuSkuGetAPIRequest, session string) (*servicecenter.TaobaoFuwuSkuGetAPIResponse, error) {
    var resp servicecenter.TaobaoFuwuSkuGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
