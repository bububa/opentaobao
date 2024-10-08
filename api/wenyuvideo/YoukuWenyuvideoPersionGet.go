package wenyuvideo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wenyuvideo"
)

// YoukuWenyuvideoPersionGet 根据优酷人物ID获取人物详情页，包含相关影视和相关人物
// youku.wenyuvideo.persion.get
//
// 根据优酷人物ID获取人物详情页，包含相关影视和相关人物
func YoukuWenyuvideoPersionGet(ctx context.Context, clt *core.SDKClient, req *wenyuvideo.YoukuWenyuvideoPersionGetAPIRequest, resp *wenyuvideo.YoukuWenyuvideoPersionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
