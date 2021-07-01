package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据园区ID获取楼宇 
alibaba.campus.space.building.getbycampusid

根据园区ID获取楼宇
HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
HSF方法名称：getBuildingList
*/
func AlibabaCampusSpaceBuildingGetbycampusid(clt *core.SDKClient, req *campus.AlibabaCampusSpaceBuildingGetbycampusidAPIRequest, session string) (*campus.AlibabaCampusSpaceBuildingGetbycampusidAPIResponse, error) {
    var resp campus.AlibabaCampusSpaceBuildingGetbycampusidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
