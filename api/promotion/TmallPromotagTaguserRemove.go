package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
给用户移除优惠标签 
tmall.promotag.taguser.remove

给用户载体去标
*/
func TmallPromotagTaguserRemove(clt *core.SDKClient, req *promotion.TmallPromotagTaguserRemoveRequest, session string) (*promotion.TmallPromotagTaguserRemoveResponse, error) {
    var resp promotion.TmallPromotagTaguserRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
