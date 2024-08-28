package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateRelationshipwithroomGet 查询rpId
// taobao.xhotel.rate.relationshipwithroom.get
//
// 某个卖家根据rpId查询所有的gid，可分页，不填分页信息则默认显示第一页。
func TaobaoXhotelRateRelationshipwithroomGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateRelationshipwithroomGetAPIRequest, resp *xhotelitem.TaobaoXhotelRateRelationshipwithroomGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
