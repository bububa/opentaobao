package aliexpress

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliexpress"
)

/* 
ISV更新INS私信发送的结果 
aliexpress.social.ins.directresult.update

ISV更新INS私信发送的结果
*/
func AliexpressSocialInsDirectresultUpdate(clt *core.SDKClient, req *aliexpress.AliexpressSocialInsDirectresultUpdateAPIRequest, session string) (*aliexpress.AliexpressSocialInsDirectresultUpdateAPIResponse, error) {
    var resp aliexpress.AliexpressSocialInsDirectresultUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
