package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商】——退票订单列表提取 
taobao.alitrip.seller.refundorderlist.fetch

代理商纬度退票订单列表提取
*/
func TaobaoAlitripSellerRefundorderlistFetch(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundorderlistFetchRequest, session string) (*jipiao.TaobaoAlitripSellerRefundorderlistFetchAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerRefundorderlistFetchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
