package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据类别编码查询类别 
alibaba.campus.space.type.getbycode

根据类别编码查询类别
HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
HSF方法名称：getByCode
*/
func AlibabaCampusSpaceTypeGetbycode(clt *core.SDKClient, req *campus.AlibabaCampusSpaceTypeGetbycodeRequest, session string) (*campus.AlibabaCampusSpaceTypeGetbycodeAPIResponse, error) {
    var resp campus.AlibabaCampusSpaceTypeGetbycodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
