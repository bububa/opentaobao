package fivee

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fivee"
)

/* 
查询商信息 
taobao.fivee.company.get

资质共享平台查询商信息
*/
func TaobaoFiveeCompanyGet(clt *core.SDKClient, req *fivee.TaobaoFiveeCompanyGetRequest, session string) (*fivee.TaobaoFiveeCompanyGetAPIResponse, error) {
    var resp fivee.TaobaoFiveeCompanyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
