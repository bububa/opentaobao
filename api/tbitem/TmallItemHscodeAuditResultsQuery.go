package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemHscodeAuditResultsQuery 商品hscode信息审核状态查询接口
// tmall.item.hscode.audit.results.query
//
// 通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
func TmallItemHscodeAuditResultsQuery(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemHscodeAuditResultsQueryAPIRequest, resp *tbitem.TmallItemHscodeAuditResultsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
