package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
供应商或分销商获取合作关系信息 
taobao.fenxiao.cooperation.get

获取供应商的合作关系信息
*/
func TaobaoFenxiaoCooperationGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoCooperationGetRequest, session string) (*fenxiao.TaobaoFenxiaoCooperationGetResponse, error) {
    var resp fenxiao.TaobaoFenxiaoCooperationGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
