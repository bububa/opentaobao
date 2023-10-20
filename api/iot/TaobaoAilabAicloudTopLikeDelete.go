package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtoplikedelete 取消收藏
// taobao.ailab.aicloud.top.like.delete
//
// 取消收藏
func Taobaoailabaicloudtoplikedelete(clt *core.SDKClient, req *iot.TaobaoailabaicloudtoplikedeleteAPIRequest, session string) (*iot.TaobaoailabaicloudtoplikedeleteAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtoplikedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
