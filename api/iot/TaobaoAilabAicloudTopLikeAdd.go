package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtoplikeadd 增加收藏
// taobao.ailab.aicloud.top.like.add
//
// 将制定内容加入收藏
func Taobaoailabaicloudtoplikeadd(clt *core.SDKClient, req *iot.TaobaoailabaicloudtoplikeaddAPIRequest, session string) (*iot.TaobaoailabaicloudtoplikeaddAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtoplikeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
