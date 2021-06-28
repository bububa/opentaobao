package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
手淘定向优惠打标接口 
taobao.ump.shoutaotag.add

手淘定向优惠的优惠标签打标接口
给特定的手淘买家打上优惠标记，标记承载在自己的业务标签库中，标签有效期为7天。
*/
func TaobaoUmpShoutaotagAdd(clt *core.SDKClient, req *promotion.TaobaoUmpShoutaotagAddRequest, session string) (*promotion.TaobaoUmpShoutaotagAddAPIResponse, error) {
    var resp promotion.TaobaoUmpShoutaotagAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
