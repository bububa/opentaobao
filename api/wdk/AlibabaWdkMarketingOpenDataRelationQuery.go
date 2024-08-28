package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenDataRelationQuery 数据关联关系查询
// alibaba.wdk.marketing.open.data.relation.query
//
// 数据关联关系查询
func AlibabaWdkMarketingOpenDataRelationQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenDataRelationQueryAPIRequest, resp *wdk.AlibabaWdkMarketingOpenDataRelationQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
