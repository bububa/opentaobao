package idleparttime

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleparttime"
)

// Alibabaidleparttimenotify 兼职通知接口
// alibaba.idle.parttime.notify
//
// 服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态)
func Alibabaidleparttimenotify(clt *core.SDKClient, req *idleparttime.AlibabaidleparttimenotifyAPIRequest, session string) (*idleparttime.AlibabaidleparttimenotifyAPIResponse, error) {
	var resp idleparttime.AlibabaidleparttimenotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
