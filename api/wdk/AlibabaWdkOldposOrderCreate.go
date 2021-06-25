package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
淘鲜达外部商户老pos机产生的订单同步进淘鲜达 
alibaba.wdk.oldpos.order.create

淘鲜达外部商户老pos机产生的订单同步进淘鲜达
*/
func AlibabaWdkOldposOrderCreate(clt *core.SDKClient, req *wdk.AlibabaWdkOldposOrderCreateRequest, session string) (*wdk.AlibabaWdkOldposOrderCreateResponse, error) {
    var resp wdk.AlibabaWdkOldposOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
