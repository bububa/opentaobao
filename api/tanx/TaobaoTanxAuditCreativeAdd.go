package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
创意预审新增接口 
taobao.tanx.audit.creative.add

创意预审新增接口
*/
func TaobaoTanxAuditCreativeAdd(clt *core.SDKClient, req *tanx.TaobaoTanxAuditCreativeAddRequest, session string) (*tanx.TaobaoTanxAuditCreativeAddResponse, error) {
    var resp tanx.TaobaoTanxAuditCreativeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
