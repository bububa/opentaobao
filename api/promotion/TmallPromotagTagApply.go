package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠标签申请 
tmall.promotag.tag.apply

创建优惠标签
*/
func TmallPromotagTagApply(clt *core.SDKClient, req *promotion.TmallPromotagTagApplyRequest, session string) (*promotion.TmallPromotagTagApplyAPIResponse, error) {
    var resp promotion.TmallPromotagTagApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
