package singletreasure

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityItemBatchadd 批量添加商品接口
// taobao.singletreasure.activity.item.batchadd
//
// 向活动中批量添加商品优惠
func TaobaoSingletreasureActivityItemBatchadd(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemBatchaddAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityItemBatchaddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
