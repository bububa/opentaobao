package moscm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/moscm"
)

/* 
调整库存 
alibaba.mos.goods.adjust

库存调整接口
*/
func AlibabaMosGoodsAdjust(clt *core.SDKClient, req *moscm.AlibabaMosGoodsAdjustRequest, session string) (*moscm.AlibabaMosGoodsAdjustAPIResponse, error) {
    var resp moscm.AlibabaMosGoodsAdjustAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
