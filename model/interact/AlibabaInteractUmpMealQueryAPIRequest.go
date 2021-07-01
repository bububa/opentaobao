package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractUmpMealQueryAPIRequest
淘宝卖家搭配套餐查询 API请求
alibaba.interact.ump.meal.query

查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接 */
type AlibabaInteractUmpMealQueryAPIRequest struct {
	model.Params
}

// New
