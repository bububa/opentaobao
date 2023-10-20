package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpDetailDelete 删除活动详情
// taobao.ump.detail.delete
//
// 删除活动详情
func TaobaoUmpDetailDelete(clt *core.SDKClient, req *promotion.TaobaoUmpDetailDeleteAPIRequest, resp *promotion.TaobaoUmpDetailDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
