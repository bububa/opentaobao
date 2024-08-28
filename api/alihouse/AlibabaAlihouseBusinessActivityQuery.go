package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseBusinessActivityQuery 电商券活动公司数据查询
// alibaba.alihouse.business.activity.query
//
// 电商券活动公司数据查询
func AlibabaAlihouseBusinessActivityQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseBusinessActivityQueryAPIRequest, resp *alihouse.AlibabaAlihouseBusinessActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
