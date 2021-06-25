package traderate

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/traderate"
)

/* 
飞猪通用评价接口 
taobao.fliggy.wrate.getmixratelist

飞猪评价通用接口
*/
func TaobaoFliggyWrateGetmixratelist(clt *core.SDKClient, req *traderate.TaobaoFliggyWrateGetmixratelistRequest, session string) (*traderate.TaobaoFliggyWrateGetmixratelistResponse, error) {
    var resp traderate.TaobaoFliggyWrateGetmixratelistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
