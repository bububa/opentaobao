package idleparttime

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleparttime"
)

// Alibabaidleparttimesynclog 兼职同步日志
// alibaba.idle.parttime.synclog
//
// 提供给供应商查询的接口
func Alibabaidleparttimesynclog(clt *core.SDKClient, req *idleparttime.AlibabaidleparttimesynclogAPIRequest, session string) (*idleparttime.AlibabaidleparttimesynclogAPIResponse, error) {
	var resp idleparttime.AlibabaidleparttimesynclogAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
