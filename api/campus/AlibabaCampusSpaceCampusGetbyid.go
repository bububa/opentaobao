package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据园区id获取园区信息 
alibaba.campus.space.campus.getbyid

根据园区id获取园区信息
HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
HSF方法名称：getCampusById
*/
func AlibabaCampusSpaceCampusGetbyid(clt *core.SDKClient, req *campus.AlibabaCampusSpaceCampusGetbyidAPIRequest, session string) (*campus.AlibabaCampusSpaceCampusGetbyidAPIResponse, error) {
    var resp campus.AlibabaCampusSpaceCampusGetbyidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
