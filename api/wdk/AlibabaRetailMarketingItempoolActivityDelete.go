package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
删除商品池活动【同城零售】 
alibaba.retail.marketing.itempool.activity.delete

同城零售商品池活动删除
*/
func AlibabaRetailMarketingItempoolActivityDelete(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityDeleteAPIRequest, session string) (*wdk.AlibabaRetailMarketingItempoolActivityDeleteAPIResponse, error) {
    var resp wdk.AlibabaRetailMarketingItempoolActivityDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
