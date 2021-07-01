package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
饿了么日维度对账单查询 
alibaba.wdk.eleme.bill.get

查询饿了么日维度对账单信息
*/
func AlibabaWdkElemeBillGet(clt *core.SDKClient, req *wdk.AlibabaWdkElemeBillGetAPIRequest, session string) (*wdk.AlibabaWdkElemeBillGetAPIResponse, error) {
    var resp wdk.AlibabaWdkElemeBillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
