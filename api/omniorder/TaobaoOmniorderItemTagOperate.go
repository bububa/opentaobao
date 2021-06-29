package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
全渠道商品打标与去标 
taobao.omniorder.item.tag.operate

用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
*/
func TaobaoOmniorderItemTagOperate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderItemTagOperateRequest, session string) (*omniorder.TaobaoOmniorderItemTagOperateAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderItemTagOperateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
