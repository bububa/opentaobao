package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
增加商品池里面的类目 
alibaba.wdk.marketing.itempool.addcategory

增加商品池里面的类目
*/
func AlibabaWdkMarketingItempoolAddcategory(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolAddcategoryRequest, session string) (*wdk.AlibabaWdkMarketingItempoolAddcategoryResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolAddcategoryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
