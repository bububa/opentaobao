package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
新增产品线 
taobao.fenxiao.productcat.add

新增产品线
*/
func TaobaoFenxiaoProductcatAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatAddRequest, session string) (*fenxiao.TaobaoFenxiaoProductcatAddResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductcatAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
