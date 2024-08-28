package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripItemSchemaUpdate 使用schema进行商品编辑
// alitrip.item.schema.update
//
// 飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
func AlitripItemSchemaUpdate(ctx context.Context, clt *core.SDKClient, req *travel.AlitripItemSchemaUpdateAPIRequest, resp *travel.AlitripItemSchemaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
