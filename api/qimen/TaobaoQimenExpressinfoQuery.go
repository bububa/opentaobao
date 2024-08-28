package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenExpressinfoQuery 配送公司信息查询接口
// taobao.qimen.expressinfo.query
//
// 配送公司信息查询
func TaobaoQimenExpressinfoQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenExpressinfoQueryAPIRequest, resp *qimen.TaobaoQimenExpressinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
