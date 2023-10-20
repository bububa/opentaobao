package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpActivitiesListGet 营销活动列表查询
// taobao.ump.activities.list.get
//
// 按照营销活动id的列表ids，查询对应的营销活动列表。
func TaobaoUmpActivitiesListGet(clt *core.SDKClient, req *promotion.TaobaoUmpActivitiesListGetAPIRequest, resp *promotion.TaobaoUmpActivitiesListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
