package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopMessageList 获取留言列表
// taobao.ailab.aicloud.top.message.list
//
// 根据指定参数获取留言列表
func TaobaoAilabAicloudTopMessageList(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageListAPIRequest, session string) (*iot.TaobaoAilabAicloudTopMessageListAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopMessageListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
