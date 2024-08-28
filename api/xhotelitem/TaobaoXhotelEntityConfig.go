package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelEntityConfig 飞猪商品各实体通用配置
// taobao.xhotel.entity.config
//
// 飞猪商品各实体通用配置服务
func TaobaoXhotelEntityConfig(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelEntityConfigAPIRequest, resp *xhotelitem.TaobaoXhotelEntityConfigAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
