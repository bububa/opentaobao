package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
库存初始化 
taobao.inventory.initial

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。
*/
func TaobaoInventoryInitial(clt *core.SDKClient, req *fenxiao.TaobaoInventoryInitialRequest, session string) (*fenxiao.TaobaoInventoryInitialAPIResponse, error) {
    var resp fenxiao.TaobaoInventoryInitialAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
