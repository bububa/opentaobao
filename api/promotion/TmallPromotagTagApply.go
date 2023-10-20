package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Tmallpromotagtagapply 优惠标签申请
// tmall.promotag.tag.apply
//
// 创建优惠标签
func Tmallpromotagtagapply(clt *core.SDKClient, req *promotion.TmallpromotagtagapplyAPIRequest, session string) (*promotion.TmallpromotagtagapplyAPIResponse, error) {
	var resp promotion.TmallpromotagtagapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
