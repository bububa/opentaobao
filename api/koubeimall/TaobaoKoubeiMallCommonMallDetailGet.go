package koubeimall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/koubeimall"
)

/* 
查询商圈详细信息 
taobao.koubei.mall.common.mall.detail.get

查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
*/
func TaobaoKoubeiMallCommonMallDetailGet(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonMallDetailGetRequest, session string) (*koubeimall.TaobaoKoubeiMallCommonMallDetailGetAPIResponse, error) {
    var resp koubeimall.TaobaoKoubeiMallCommonMallDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
