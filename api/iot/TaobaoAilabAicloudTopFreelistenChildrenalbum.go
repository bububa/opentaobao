package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopfreelistenchildrenalbum 儿童音频列表
// taobao.ailab.aicloud.top.freelisten.childrenalbum
//
// 儿童音频列表
func Taobaoailabaicloudtopfreelistenchildrenalbum(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopfreelistenchildrenalbumAPIRequest, session string) (*iot.TaobaoailabaicloudtopfreelistenchildrenalbumAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopfreelistenchildrenalbumAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
