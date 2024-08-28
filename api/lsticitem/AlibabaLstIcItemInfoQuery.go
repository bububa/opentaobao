package lsticitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsticitem"
)

// AlibabaLstIcItemInfoQuery 商品信息查询
// alibaba.lst.ic.item.info.query
//
// 查询商品信息
func AlibabaLstIcItemInfoQuery(ctx context.Context, clt *core.SDKClient, req *lsticitem.AlibabaLstIcItemInfoQueryAPIRequest, resp *lsticitem.AlibabaLstIcItemInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
