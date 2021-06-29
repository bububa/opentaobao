package koubeimall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/koubeimall"
)

/* 
查询门店推荐菜信息 
taobao.koubei.mall.common.store.display.goods.list

提供查询口碑商圈内的门店推荐菜信息
*/
func TaobaoKoubeiMallCommonStoreDisplayGoodsList(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStoreDisplayGoodsListRequest, session string) (*koubeimall.TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse, error) {
    var resp koubeimall.TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
