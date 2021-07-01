package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据分组ID查询相应的空间单元 
alibaba.campus.space.unit.getlistbygroupid

根据分组ID查询相应的空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByGroupId
*/
func AlibabaCampusSpaceUnitGetlistbygroupid(clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest, session string) (*campus.AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse, error) {
    var resp campus.AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
