package retail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/retail"
)

/* 
查询货架和位置数据 
alibaba.interact.retail.queryshelflocation

查询货架和位置数据
*/
func AlibabaInteractRetailQueryshelflocation(clt *core.SDKClient, req *retail.AlibabaInteractRetailQueryshelflocationRequest, session string) (*retail.AlibabaInteractRetailQueryshelflocationAPIResponse, error) {
    var resp retail.AlibabaInteractRetailQueryshelflocationAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
