package ascpqcc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpqcc"
)

/* 
魅力惠样品解除父子商品关系 
alibaba.ascp.qcc.sample.cancel.item.relation

品控中心魅力惠样品解除父子商品关系
*/
func AlibabaAscpQccSampleCancelItemRelation(clt *core.SDKClient, req *ascpqcc.AlibabaAscpQccSampleCancelItemRelationRequest, session string) (*ascpqcc.AlibabaAscpQccSampleCancelItemRelationAPIResponse, error) {
    var resp ascpqcc.AlibabaAscpQccSampleCancelItemRelationAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
