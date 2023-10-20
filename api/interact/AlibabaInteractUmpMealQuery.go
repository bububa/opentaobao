package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractumpmealquery 淘宝卖家搭配套餐查询
// alibaba.interact.ump.meal.query
//
// 查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
func Alibabainteractumpmealquery(clt *core.SDKClient, req *interact.AlibabainteractumpmealqueryAPIRequest, session string) (*interact.AlibabainteractumpmealqueryAPIResponse, error) {
	var resp interact.AlibabainteractumpmealqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
