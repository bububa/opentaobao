package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* TaobaoAilabAicloudTopFeedlistGet
获取对话流列表
taobao.ailab.aicloud.top.feedlist.get

获取指定应用的对话流信息 */
func TaobaoAilabAicloudTopFeedlistGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopFeedlistGetAPIRequest, session string) (*iot.TaobaoAilabAicloudTopFeedlistGetAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopFeedlistGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
