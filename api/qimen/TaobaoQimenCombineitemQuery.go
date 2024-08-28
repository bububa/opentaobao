package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenCombineitemQuery 组合货品关系查询接口
// taobao.qimen.combineitem.query
//
// 组合货品关系查询
func TaobaoQimenCombineitemQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenCombineitemQueryAPIRequest, resp *qimen.TaobaoQimenCombineitemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
