package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
第三方车型数据上传 
tmall.aliauto.wisdomdata.model.recive

天猫汽车对外提供的汽车车型数据上传接口
*/
func TmallAliautoWisdomdataModelRecive(clt *core.SDKClient, req *tmallcar.TmallAliautoWisdomdataModelReciveAPIRequest, session string) (*tmallcar.TmallAliautoWisdomdataModelReciveAPIResponse, error) {
    var resp tmallcar.TmallAliautoWisdomdataModelReciveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
