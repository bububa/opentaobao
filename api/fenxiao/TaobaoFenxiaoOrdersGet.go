package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询采购单信息 
taobao.fenxiao.orders.get

分销商或供应商均可用此接口查询采购单信息（代销）； (发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
*/
func TaobaoFenxiaoOrdersGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrdersGetRequest, session string) (*fenxiao.TaobaoFenxiaoOrdersGetResponse, error) {
    var resp fenxiao.TaobaoFenxiaoOrdersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
