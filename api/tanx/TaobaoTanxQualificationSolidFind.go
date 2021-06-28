package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
客户固态共享资质 
taobao.tanx.qualification.solid.find

接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。
*/
func TaobaoTanxQualificationSolidFind(clt *core.SDKClient, req *tanx.TaobaoTanxQualificationSolidFindRequest, session string) (*tanx.TaobaoTanxQualificationSolidFindAPIResponse, error) {
    var resp tanx.TaobaoTanxQualificationSolidFindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
