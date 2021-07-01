package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
线下预存款流水增加 
taobao.fenxiao.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加）
*/
func TaobaoFenxiaoTradePrepayOfflineAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoTradePrepayOfflineAddAPIRequest, session string) (*fenxiao.TaobaoFenxiaoTradePrepayOfflineAddAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoTradePrepayOfflineAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
