package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpActivityGet 查询营销活动
// taobao.ump.activity.get
//
// 查询营销活动
func TaobaoUmpActivityGet(clt *core.SDKClient, req *promotion.TaobaoUmpActivityGetAPIRequest, resp *promotion.TaobaoUmpActivityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
