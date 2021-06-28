package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询商品池活动下的商品 
alibaba.wdk.marketing.itempool.queryitems

查询商品池活动下面的商品
*/
func AlibabaWdkMarketingItempoolQueryitems(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolQueryitemsRequest, session string) (*wdk.AlibabaWdkMarketingItempoolQueryitemsAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolQueryitemsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
