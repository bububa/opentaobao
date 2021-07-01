package koubeimall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/koubeimall"
)

/* 
查询商品详情信息 
taobao.koubei.mall.common.item.detail.query

查询口碑综合体内商品详情信息
*/
func TaobaoKoubeiMallCommonItemDetailQuery(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonItemDetailQueryAPIRequest, session string) (*koubeimall.TaobaoKoubeiMallCommonItemDetailQueryAPIResponse, error) {
    var resp koubeimall.TaobaoKoubeiMallCommonItemDetailQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
