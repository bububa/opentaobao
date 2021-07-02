package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotagTagFind 查询标签接口
// tmall.promotag.tag.find
//
// 查询用户创建的所有标签
func TmallPromotagTagFind(clt *core.SDKClient, req *promotion.TmallPromotagTagFindAPIRequest, session string) (*promotion.TmallPromotagTagFindAPIResponse, error) {
	var resp promotion.TmallPromotagTagFindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
