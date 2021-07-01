package koubeimall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/koubeimall"
)

/* 
分页查询综合体内的门店列表信息 
taobao.koubei.mall.common.store.page

分页查询综合体内的门店列表信息
*/
func TaobaoKoubeiMallCommonStorePage(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStorePageAPIRequest, session string) (*koubeimall.TaobaoKoubeiMallCommonStorePageAPIResponse, error) {
    var resp koubeimall.TaobaoKoubeiMallCommonStorePageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
