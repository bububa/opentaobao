package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopLikeFilter 过滤列表歌曲存在于收藏列表的
// taobao.ailab.aicloud.top.like.filter
//
// 过滤出传入列表歌曲存在于收藏列表的
func TaobaoAilabAicloudTopLikeFilter(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeFilterAPIRequest, resp *iot.TaobaoAilabAicloudTopLikeFilterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
