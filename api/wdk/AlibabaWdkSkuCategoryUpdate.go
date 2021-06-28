package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商家类目修改接口 
alibaba.wdk.sku.category.update

商家类目修改接口
*/
func AlibabaWdkSkuCategoryUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryUpdateRequest, session string) (*wdk.AlibabaWdkSkuCategoryUpdateAPIResponse, error) {
    var resp wdk.AlibabaWdkSkuCategoryUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
