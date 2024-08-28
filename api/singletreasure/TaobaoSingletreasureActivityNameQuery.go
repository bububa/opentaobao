package singletreasure

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityNameQuery 查询官方的活动名称接口
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
func TaobaoSingletreasureActivityNameQuery(ctx context.Context, clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityNameQueryAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityNameQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
