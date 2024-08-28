package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenItemmappingCreate 前后端商品映射接口
// taobao.qimen.itemmapping.create
//
// 前后端商品映射
func TaobaoQimenItemmappingCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenItemmappingCreateAPIRequest, resp *qimen.TaobaoQimenItemmappingCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
