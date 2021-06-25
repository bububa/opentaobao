package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
获取分销商信息 
taobao.fenxiao.distributors.get

查询和当前登录供应商有合作关系的分销商的信息
*/
func TaobaoFenxiaoDistributorsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorsGetRequest, session string) (*fenxiao.TaobaoFenxiaoDistributorsGetResponse, error) {
    var resp fenxiao.TaobaoFenxiaoDistributorsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
