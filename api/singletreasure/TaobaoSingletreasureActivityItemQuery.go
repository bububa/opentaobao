package singletreasure

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityItemQuery 查询活动下的优惠信息
// taobao.singletreasure.activity.item.query
//
// 分页查询活动下的商品优惠信息
func TaobaoSingletreasureActivityItemQuery(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemQueryAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
