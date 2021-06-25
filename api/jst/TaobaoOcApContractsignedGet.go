package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
用户是否签署支付宝代扣协议 
taobao.oc.ap.contractsigned.get

用户是否签署支付宝代扣协议
*/
func TaobaoOcApContractsignedGet(clt *core.SDKClient, req *jst.TaobaoOcApContractsignedGetRequest, session string) (*jst.TaobaoOcApContractsignedGetResponse, error) {
    var resp jst.TaobaoOcApContractsignedGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
