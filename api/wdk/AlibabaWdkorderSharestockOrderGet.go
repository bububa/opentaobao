package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
猫超商户订单拉取 
alibaba.wdkorder.sharestock.order.get

商户拉取猫超订单数据
*/
func AlibabaWdkorderSharestockOrderGet(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockOrderGetRequest, session string) (*wdk.AlibabaWdkorderSharestockOrderGetResponse, error) {
    var resp wdk.AlibabaWdkorderSharestockOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
