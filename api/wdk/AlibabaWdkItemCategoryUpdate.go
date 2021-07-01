package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
修改类目 
alibaba.wdk.item.category.update

修改类目
*/
func AlibabaWdkItemCategoryUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkItemCategoryUpdateAPIRequest, session string) (*wdk.AlibabaWdkItemCategoryUpdateAPIResponse, error) {
    var resp wdk.AlibabaWdkItemCategoryUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
