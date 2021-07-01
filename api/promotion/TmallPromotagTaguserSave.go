package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
给用户打上优惠标签 
tmall.promotag.taguser.save

给用户载体打标
*/
func TmallPromotagTaguserSave(clt *core.SDKClient, req *promotion.TmallPromotagTaguserSaveAPIRequest, session string) (*promotion.TmallPromotagTaguserSaveAPIResponse, error) {
    var resp promotion.TmallPromotagTaguserSaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
