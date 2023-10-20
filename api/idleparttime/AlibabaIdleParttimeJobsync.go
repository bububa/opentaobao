package idleparttime

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleparttime"
)

// Alibabaidleparttimejobsync 兼职岗位同步
// alibaba.idle.parttime.jobsync
//
// 服务商同步岗位信息给闲鱼
func Alibabaidleparttimejobsync(clt *core.SDKClient, req *idleparttime.AlibabaidleparttimejobsyncAPIRequest, session string) (*idleparttime.AlibabaidleparttimejobsyncAPIResponse, error) {
	var resp idleparttime.AlibabaidleparttimejobsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
