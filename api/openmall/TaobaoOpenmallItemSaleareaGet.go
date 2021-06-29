package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
查询商品可售区域 
taobao.openmall.item.salearea.get

获取openmall商品的可售区域
*/
func TaobaoOpenmallItemSaleareaGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallItemSaleareaGetRequest, session string) (*openmall.TaobaoOpenmallItemSaleareaGetAPIResponse, error) {
    var resp openmall.TaobaoOpenmallItemSaleareaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
