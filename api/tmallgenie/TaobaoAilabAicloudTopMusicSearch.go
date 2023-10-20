package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMusicSearch 对外音乐搜索服务
// taobao.ailab.aicloud.top.music.search
//
// 供厂商获取音乐列表
func TaobaoAilabAicloudTopMusicSearch(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMusicSearchAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMusicSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
