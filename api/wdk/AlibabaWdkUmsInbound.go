package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
入库-ERP下发单 
alibaba.wdk.ums.inbound

入库-ERP下发单
*/
func AlibabaWdkUmsInbound(clt *core.SDKClient, req *wdk.AlibabaWdkUmsInboundAPIRequest, session string) (*wdk.AlibabaWdkUmsInboundAPIResponse, error) {
    var resp wdk.AlibabaWdkUmsInboundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
