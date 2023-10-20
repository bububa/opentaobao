package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscmjsactivitydelete 删除满就送活动
// taobao.promotionmisc.mjs.activity.delete
//
// 删除满就送活动
func Taobaopromotionmiscmjsactivitydelete(clt *core.SDKClient, req *promotion.TaobaopromotionmiscmjsactivitydeleteAPIRequest, session string) (*promotion.TaobaopromotionmiscmjsactivitydeleteAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscmjsactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
