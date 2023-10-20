package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpActivityUpdate 修改活动信息
// taobao.ump.activity.update
//
// 修改营销活动
func TaobaoUmpActivityUpdate(clt *core.SDKClient, req *promotion.TaobaoUmpActivityUpdateAPIRequest, resp *promotion.TaobaoUmpActivityUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
