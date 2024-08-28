package singletreasure

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityUpdate 修改活动接口
// taobao.singletreasure.activity.update
//
// 修改活动接口
func TaobaoSingletreasureActivityUpdate(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityUpdateAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
