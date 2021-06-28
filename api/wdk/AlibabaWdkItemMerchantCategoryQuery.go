package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询商品的商家叶子类目 
alibaba.wdk.item.merchant.category.query

查询商品的商家叶子类目
*/
func AlibabaWdkItemMerchantCategoryQuery(clt *core.SDKClient, req *wdk.AlibabaWdkItemMerchantCategoryQueryRequest, session string) (*wdk.AlibabaWdkItemMerchantCategoryQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkItemMerchantCategoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
