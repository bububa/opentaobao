package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateRelationshipwithrpGet 根据gid查询卖家下所有的rpId
// taobao.xhotel.rate.relationshipwithrp.get
//
// 根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据
func TaobaoXhotelRateRelationshipwithrpGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateRelationshipwithrpGetAPIRequest, resp *xhotelitem.TaobaoXhotelRateRelationshipwithrpGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
