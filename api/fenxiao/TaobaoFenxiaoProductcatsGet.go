package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询产品线列表 
taobao.fenxiao.productcats.get

查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
*/
func TaobaoFenxiaoProductcatsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatsGetRequest, session string) (*fenxiao.TaobaoFenxiaoProductcatsGetResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductcatsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
