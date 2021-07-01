package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
空间单元列表带业务属性实例 
alibaba.campus.space.unit.getspaceunitlistwithattr

空间单元列表带业务属性实例
*/
func AlibabaCampusSpaceUnitGetspaceunitlistwithattr(clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest, session string) (*campus.AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse, error) {
    var resp campus.AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
