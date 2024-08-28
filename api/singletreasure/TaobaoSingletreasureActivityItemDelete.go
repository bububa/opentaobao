package singletreasure

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityItemDelete 删除单品优惠接口
// taobao.singletreasure.activity.item.delete
//
// 删除单品优惠接口
func TaobaoSingletreasureActivityItemDelete(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemDeleteAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityItemDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
