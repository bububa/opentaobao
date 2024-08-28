package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// TaobaoAilabAicloudTopHotwordsUpdate 更新热词
// taobao.ailab.aicloud.top.hotwords.update
//
// 更新ASR热词
func TaobaoAilabAicloudTopHotwordsUpdate(ctx context.Context, clt *core.SDKClient, req *alilabs.TaobaoAilabAicloudTopHotwordsUpdateAPIRequest, resp *alilabs.TaobaoAilabAicloudTopHotwordsUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
