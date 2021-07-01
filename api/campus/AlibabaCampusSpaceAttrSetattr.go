package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
新增业务属性实例接口 
alibaba.campus.space.attr.setattr

新增业务属性实例接口
*/
func AlibabaCampusSpaceAttrSetattr(clt *core.SDKClient, req *campus.AlibabaCampusSpaceAttrSetattrAPIRequest, session string) (*campus.AlibabaCampusSpaceAttrSetattrAPIResponse, error) {
    var resp campus.AlibabaCampusSpaceAttrSetattrAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
