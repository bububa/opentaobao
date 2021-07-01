package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据园区id及TypeId获取空间分组 
alibaba.campus.space.group.getlistbycampusandtype

根据园区id及TypeId获取空间分组
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getListByCampusAndType
*/
func AlibabaCampusSpaceGroupGetlistbycampusandtype(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest, session string) (*campus.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse, error) {
    var resp campus.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
