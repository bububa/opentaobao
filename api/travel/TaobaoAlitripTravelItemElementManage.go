package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
【API3.0】资源元素管理接口 
taobao.alitrip.travel.item.element.manage

资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享
*/
func TaobaoAlitripTravelItemElementManage(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemElementManageRequest, session string) (*travel.TaobaoAlitripTravelItemElementManageResponse, error) {
    var resp travel.TaobaoAlitripTravelItemElementManageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
