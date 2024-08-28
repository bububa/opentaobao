package singletreasure

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityItemUpdate 更新单品优惠接口
// taobao.singletreasure.activity.item.update
//
// 更新单品优惠接口
func TaobaoSingletreasureActivityItemUpdate(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemUpdateAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityItemUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
