package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// TaobaoAilabAicloudTopHotwordsGet 获取热词
// taobao.ailab.aicloud.top.hotwords.get
//
// 获取ASR热词
func TaobaoAilabAicloudTopHotwordsGet(ctx context.Context, clt *core.SDKClient, req *alilabs.TaobaoAilabAicloudTopHotwordsGetAPIRequest, resp *alilabs.TaobaoAilabAicloudTopHotwordsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
