package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
渠道分销供应商上传线下流水预存款（减少） 
taobao.fenxiao.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少）
*/
func TaobaoFenxiaoTradePrepayOfflineReduce(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoTradePrepayOfflineReduceRequest, session string) (*fenxiao.TaobaoFenxiaoTradePrepayOfflineReduceResponse, error) {
    var resp fenxiao.TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
