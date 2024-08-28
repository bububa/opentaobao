package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenItemmappingQuery 前后端商品映射查询接口
// taobao.qimen.itemmapping.query
//
// 前后端商品映射查询接口
func TaobaoQimenItemmappingQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenItemmappingQueryAPIRequest, resp *qimen.TaobaoQimenItemmappingQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
