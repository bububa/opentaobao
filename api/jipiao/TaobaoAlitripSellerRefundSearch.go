package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商】退票申请单列表 
taobao.alitrip.seller.refund.search

查询退票申请单列表
*/
func TaobaoAlitripSellerRefundSearch(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundSearchRequest, session string) (*jipiao.TaobaoAlitripSellerRefundSearchAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerRefundSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
