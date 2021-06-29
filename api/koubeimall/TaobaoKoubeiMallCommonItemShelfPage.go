package koubeimall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/koubeimall"
)

/* 
门店货架商品列表信息查询 
taobao.koubei.mall.common.item.shelf.page

查询口碑综合体内门店货架商品列表信息接口
*/
func TaobaoKoubeiMallCommonItemShelfPage(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonItemShelfPageRequest, session string) (*koubeimall.TaobaoKoubeiMallCommonItemShelfPageAPIResponse, error) {
    var resp koubeimall.TaobaoKoubeiMallCommonItemShelfPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
