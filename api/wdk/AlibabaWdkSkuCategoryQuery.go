package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商家类目获取接口 
alibaba.wdk.sku.category.query

商家类目获取接口
*/
func AlibabaWdkSkuCategoryQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryQueryAPIRequest, session string) (*wdk.AlibabaWdkSkuCategoryQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkSkuCategoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
