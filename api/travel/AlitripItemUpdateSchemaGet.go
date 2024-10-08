package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripItemUpdateSchemaGet 获取编辑商品的schema模板
// alitrip.item.update.schema.get
//
// 获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
func AlitripItemUpdateSchemaGet(ctx context.Context, clt *core.SDKClient, req *travel.AlitripItemUpdateSchemaGetAPIRequest, resp *travel.AlitripItemUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
