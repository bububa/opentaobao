package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmusicsearch 对外音乐搜索服务
// taobao.ailab.aicloud.top.music.search
//
// 供厂商获取音乐列表
func Taobaoailabaicloudtopmusicsearch(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmusicsearchAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmusicsearchAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmusicsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
