package alitripdivisions

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripdivisions"
)

/* 
根据中文名称与行政区划级别查询行政区划数据 
alitrip.platform.divisions.getdivisionbyname

根据中文名称与行政区划级别查询行政区划数据
*/
func AlitripPlatformDivisionsGetdivisionbyname(clt *core.SDKClient, req *alitripdivisions.AlitripPlatformDivisionsGetdivisionbynameRequest, session string) (*alitripdivisions.AlitripPlatformDivisionsGetdivisionbynameAPIResponse, error) {
    var resp alitripdivisions.AlitripPlatformDivisionsGetdivisionbynameAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
