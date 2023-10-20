package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpActivitiesGet 查询活动列表
// taobao.ump.activities.get
//
// 查询活动列表
func TaobaoUmpActivitiesGet(clt *core.SDKClient, req *promotion.TaobaoUmpActivitiesGetAPIRequest, resp *promotion.TaobaoUmpActivitiesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
