package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商】退票申请单详情 
taobao.alitrip.seller.refund.get

查询退票申请单详情
*/
func TaobaoAlitripSellerRefundGet(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundGetRequest, session string) (*jipiao.TaobaoAlitripSellerRefundGetAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerRefundGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
