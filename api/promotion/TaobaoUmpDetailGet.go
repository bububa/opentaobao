package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpDetailGet 查询活动详情
// taobao.ump.detail.get
//
// 查询活动详情
func TaobaoUmpDetailGet(clt *core.SDKClient, req *promotion.TaobaoUmpDetailGetAPIRequest, resp *promotion.TaobaoUmpDetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
