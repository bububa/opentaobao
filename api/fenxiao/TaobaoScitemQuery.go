package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询后端商品 
taobao.scitem.query

查询后端商品
*/
func TaobaoScitemQuery(clt *core.SDKClient, req *fenxiao.TaobaoScitemQueryRequest, session string) (*fenxiao.TaobaoScitemQueryResponse, error) {
    var resp fenxiao.TaobaoScitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
