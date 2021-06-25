package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
增值服务订购服务验证 
taobao.vas.service.validate

增值服务订购服务验证
*/
func TaobaoVasServiceValidate(clt *core.SDKClient, req *servicecenter.TaobaoVasServiceValidateRequest, session string) (*servicecenter.TaobaoVasServiceValidateResponse, error) {
    var resp servicecenter.TaobaoVasServiceValidateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
