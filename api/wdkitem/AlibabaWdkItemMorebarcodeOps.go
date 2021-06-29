package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
商品一品多码维护操作 
alibaba.wdk.item.morebarcode.ops

商品一品多码维护操作
*/
func AlibabaWdkItemMorebarcodeOps(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMorebarcodeOpsRequest, session string) (*wdkitem.AlibabaWdkItemMorebarcodeOpsAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemMorebarcodeOpsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
