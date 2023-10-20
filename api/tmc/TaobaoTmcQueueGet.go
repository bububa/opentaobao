package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcqueueget 获取消息队列积压情况
// taobao.tmc.queue.get
//
// 根据appkey和groupName获取消息队列积压情况
func Taobaotmcqueueget(clt *core.SDKClient, req *tmc.TaobaotmcqueuegetAPIRequest, session string) (*tmc.TaobaotmcqueuegetAPIResponse, error) {
	var resp tmc.TaobaotmcqueuegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
