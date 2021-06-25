package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
删除产品线 
taobao.fenxiao.productcat.delete

删除产品线
*/
func TaobaoFenxiaoProductcatDelete(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatDeleteRequest, session string) (*fenxiao.TaobaoFenxiaoProductcatDeleteResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductcatDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
