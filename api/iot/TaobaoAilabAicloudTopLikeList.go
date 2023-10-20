package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtoplikelist 列出收藏列表
// taobao.ailab.aicloud.top.like.list
//
// 列出收藏列表
func Taobaoailabaicloudtoplikelist(clt *core.SDKClient, req *iot.TaobaoailabaicloudtoplikelistAPIRequest, session string) (*iot.TaobaoailabaicloudtoplikelistAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtoplikelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
