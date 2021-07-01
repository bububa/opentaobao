package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* TaobaoAilabAicloudTopLikeList
列出收藏列表
taobao.ailab.aicloud.top.like.list

列出收藏列表 */
func TaobaoAilabAicloudTopLikeList(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeListAPIRequest, session string) (*iot.TaobaoAilabAicloudTopLikeListAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopLikeListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
