package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商品池排除商品【品类优惠使用】 
alibaba.wdk.marketing.itempool.excludeskucode

品类优惠新增排除池
*/
func AlibabaWdkMarketingItempoolExcludeskucode(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolExcludeskucodeRequest, session string) (*wdk.AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
