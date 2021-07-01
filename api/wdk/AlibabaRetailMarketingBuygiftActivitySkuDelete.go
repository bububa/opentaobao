package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
删除单品买赠活动商品 
alibaba.retail.marketing.buygift.activity.sku.delete

删除单品买赠活动商品信息【同城零售】
*/
func AlibabaRetailMarketingBuygiftActivitySkuDelete(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse, error) {
    var resp wdk.AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
