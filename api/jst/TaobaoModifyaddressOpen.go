package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
淘宝自助修改地址服务开通 
taobao.modifyaddress.open

商家自助修改地址服务开通
*/
func TaobaoModifyaddressOpen(clt *core.SDKClient, req *jst.TaobaoModifyaddressOpenRequest, session string) (*jst.TaobaoModifyaddressOpenResponse, error) {
    var resp jst.TaobaoModifyaddressOpenAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
