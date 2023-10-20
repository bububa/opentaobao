package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpRangeGet 查询活动范围
// taobao.ump.range.get
//
// 查询某个卖家所有参加或者不参加某项活动的物品
func TaobaoUmpRangeGet(clt *core.SDKClient, req *promotion.TaobaoUmpRangeGetAPIRequest, resp *promotion.TaobaoUmpRangeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
