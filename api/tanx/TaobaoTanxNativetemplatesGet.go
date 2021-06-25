package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
批量获取本地模板信息 
taobao.tanx.nativetemplates.get

根据传入的本地模板ID批量返回本地模板
*/
func TaobaoTanxNativetemplatesGet(clt *core.SDKClient, req *tanx.TaobaoTanxNativetemplatesGetRequest, session string) (*tanx.TaobaoTanxNativetemplatesGetResponse, error) {
    var resp tanx.TaobaoTanxNativetemplatesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
