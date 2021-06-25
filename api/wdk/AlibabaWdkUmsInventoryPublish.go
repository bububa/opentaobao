package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
初始化覆盖实物库存 
alibaba.wdk.ums.inventory.publish

先去库存这边查询当前实物库存有多少的量，然后算出来需要增加的量。接下来调用ums原来的入库语义的接口进行库存的增量补充
*/
func AlibabaWdkUmsInventoryPublish(clt *core.SDKClient, req *wdk.AlibabaWdkUmsInventoryPublishRequest, session string) (*wdk.AlibabaWdkUmsInventoryPublishResponse, error) {
    var resp wdk.AlibabaWdkUmsInventoryPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
