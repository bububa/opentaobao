package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
合作申请查询 
taobao.fenxiao.requisitions.get

合作申请查询
*/
func TaobaoFenxiaoRequisitionsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoRequisitionsGetRequest, session string) (*fenxiao.TaobaoFenxiaoRequisitionsGetResponse, error) {
    var resp fenxiao.TaobaoFenxiaoRequisitionsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
