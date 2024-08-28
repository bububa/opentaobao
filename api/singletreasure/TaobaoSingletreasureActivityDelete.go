package singletreasure

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityDelete 删除活动接口
// taobao.singletreasure.activity.delete
//
// 删除优惠活动
func TaobaoSingletreasureActivityDelete(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityDeleteAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
