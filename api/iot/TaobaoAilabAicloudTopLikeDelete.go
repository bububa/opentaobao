package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* TaobaoAilabAicloudTopLikeDelete
取消收藏
taobao.ailab.aicloud.top.like.delete

取消收藏 */
func TaobaoAilabAicloudTopLikeDelete(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeDeleteAPIRequest, session string) (*iot.TaobaoAilabAicloudTopLikeDeleteAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopLikeDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
