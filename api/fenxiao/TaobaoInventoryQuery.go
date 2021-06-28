package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询商品库存信息 
taobao.inventory.query

建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
商家查询商品总体库存信息
*/
func TaobaoInventoryQuery(clt *core.SDKClient, req *fenxiao.TaobaoInventoryQueryRequest, session string) (*fenxiao.TaobaoInventoryQueryAPIResponse, error) {
    var resp fenxiao.TaobaoInventoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
