package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripItemAddSchemaGet 获取商品发布模板
// alitrip.item.add.schema.get
//
// 发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
func AlitripItemAddSchemaGet(ctx context.Context, clt *core.SDKClient, req *travel.AlitripItemAddSchemaGetAPIRequest, resp *travel.AlitripItemAddSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
