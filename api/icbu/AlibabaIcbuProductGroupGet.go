package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductGroupGet 分组信息获取
// alibaba.icbu.product.group.get
//
// 分组信息获取
func AlibabaIcbuProductGroupGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupGetAPIRequest, resp *icbu.AlibabaIcbuProductGroupGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
