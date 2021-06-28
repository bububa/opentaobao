package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
修改资质接口 
taobao.tanx.qualification.modify

对dsp上传过的资质进行修改
*/
func TaobaoTanxQualificationModify(clt *core.SDKClient, req *tanx.TaobaoTanxQualificationModifyRequest, session string) (*tanx.TaobaoTanxQualificationModifyAPIResponse, error) {
    var resp tanx.TaobaoTanxQualificationModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
