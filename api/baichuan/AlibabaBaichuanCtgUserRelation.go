package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
用户 
alibaba.baichuan.ctg.user.relation

提供给优酷查询道长和淘宝账户的绑定关系
*/
func AlibabaBaichuanCtgUserRelation(clt *core.SDKClient, req *baichuan.AlibabaBaichuanCtgUserRelationRequest, session string) (*baichuan.AlibabaBaichuanCtgUserRelationResponse, error) {
    var resp baichuan.AlibabaBaichuanCtgUserRelationAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
