package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Tmallpromotagtagusersave 给用户打上优惠标签
// tmall.promotag.taguser.save
//
// 给用户载体打标
func Tmallpromotagtagusersave(clt *core.SDKClient, req *promotion.TmallpromotagtagusersaveAPIRequest, session string) (*promotion.TmallpromotagtagusersaveAPIResponse, error) {
	var resp promotion.TmallpromotagtagusersaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
