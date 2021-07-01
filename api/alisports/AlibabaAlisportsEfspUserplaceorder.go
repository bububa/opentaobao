package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
用户完成支付同步订单 
alibaba.alisports.efsp.userplaceorder

用户完成支付同步订单
*/
func AlibabaAlisportsEfspUserplaceorder(clt *core.SDKClient, req *alisports.AlibabaAlisportsEfspUserplaceorderAPIRequest, session string) (*alisports.AlibabaAlisportsEfspUserplaceorderAPIResponse, error) {
    var resp alisports.AlibabaAlisportsEfspUserplaceorderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
