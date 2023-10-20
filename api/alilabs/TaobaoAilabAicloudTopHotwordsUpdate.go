package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// TaobaoAilabAicloudTopHotwordsUpdate 更新热词
// taobao.ailab.aicloud.top.hotwords.update
//
// 更新ASR热词
func TaobaoAilabAicloudTopHotwordsUpdate(clt *core.SDKClient, req *alilabs.TaobaoAilabAicloudTopHotwordsUpdateAPIRequest, resp *alilabs.TaobaoAilabAicloudTopHotwordsUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
