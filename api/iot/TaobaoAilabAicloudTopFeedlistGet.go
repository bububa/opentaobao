package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopfeedlistget 获取对话流列表
// taobao.ailab.aicloud.top.feedlist.get
//
// 获取指定应用的对话流信息
func Taobaoailabaicloudtopfeedlistget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopfeedlistgetAPIRequest, session string) (*iot.TaobaoailabaicloudtopfeedlistgetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopfeedlistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
