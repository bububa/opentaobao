package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
类目查询接口 
alibaba.wdk.item.category.query

类目查询接口
*/
func AlibabaWdkItemCategoryQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemCategoryQueryAPIRequest, session string) (*wdkitem.AlibabaWdkItemCategoryQueryAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemCategoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
