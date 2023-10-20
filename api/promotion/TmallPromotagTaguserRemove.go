package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotagTaguserRemove 给用户移除优惠标签
// tmall.promotag.taguser.remove
//
// 给用户载体去标
func TmallPromotagTaguserRemove(clt *core.SDKClient, req *promotion.TmallPromotagTaguserRemoveAPIRequest, resp *promotion.TmallPromotagTaguserRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
