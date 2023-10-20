package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Tmallpromotagtaguserremove 给用户移除优惠标签
// tmall.promotag.taguser.remove
//
// 给用户载体去标
func Tmallpromotagtaguserremove(clt *core.SDKClient, req *promotion.TmallpromotagtaguserremoveAPIRequest, session string) (*promotion.TmallpromotagtaguserremoveAPIResponse, error) {
	var resp promotion.TmallpromotagtaguserremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
