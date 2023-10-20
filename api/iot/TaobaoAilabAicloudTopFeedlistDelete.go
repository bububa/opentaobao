package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopfeedlistdelete 删除单条对话流信息
// taobao.ailab.aicloud.top.feedlist.delete
//
// 删除指定的某一条对话流信息
func Taobaoailabaicloudtopfeedlistdelete(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopfeedlistdeleteAPIRequest, session string) (*iot.TaobaoailabaicloudtopfeedlistdeleteAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopfeedlistdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
