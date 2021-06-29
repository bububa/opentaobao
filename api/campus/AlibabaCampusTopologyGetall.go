package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
获取管理园区的规则拓扑接口 
alibaba.campus.topology.getall

获取所属园区的所有规则拓扑图
*/
func AlibabaCampusTopologyGetall(clt *core.SDKClient, req *campus.AlibabaCampusTopologyGetallRequest, session string) (*campus.AlibabaCampusTopologyGetallAPIResponse, error) {
    var resp campus.AlibabaCampusTopologyGetallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
