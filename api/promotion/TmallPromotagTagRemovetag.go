package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
删除标签定义 
tmall.promotag.tag.removetag

用于删除标签定义，但是要确保目前该标签没有人群在使用。
*/
func TmallPromotagTagRemovetag(clt *core.SDKClient, req *promotion.TmallPromotagTagRemovetagAPIRequest, session string) (*promotion.TmallPromotagTagRemovetagAPIResponse, error) {
    var resp promotion.TmallPromotagTagRemovetagAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
