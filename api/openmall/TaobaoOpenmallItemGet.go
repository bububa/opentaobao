package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
获取商品详情物料 
taobao.openmall.item.get

获取联盟开放的openmall商品
*/
func TaobaoOpenmallItemGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallItemGetRequest, session string) (*openmall.TaobaoOpenmallItemGetAPIResponse, error) {
    var resp openmall.TaobaoOpenmallItemGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
