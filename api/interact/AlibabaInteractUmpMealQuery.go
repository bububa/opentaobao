package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractUmpMealQuery 淘宝卖家搭配套餐查询
// alibaba.interact.ump.meal.query
//
// 查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
func AlibabaInteractUmpMealQuery(clt *core.SDKClient, req *interact.AlibabaInteractUmpMealQueryAPIRequest, resp *interact.AlibabaInteractUmpMealQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
