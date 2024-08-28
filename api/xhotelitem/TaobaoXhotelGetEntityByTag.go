package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelGetEntityByTag 根据标签查询实体
// taobao.xhotel.get.entity.by.tag
//
// 根据标签查询实体
func TaobaoXhotelGetEntityByTag(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelGetEntityByTagAPIRequest, resp *xhotelitem.TaobaoXhotelGetEntityByTagAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
