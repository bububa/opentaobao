package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanchannelget 5-IBP同步渠道接口
// alibaba.tmallgenie.scp.plan.channel.get
//
// IBP同步渠道接口
func Alibabatmallgeniescpplanchannelget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanchannelgetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanchannelgetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanchannelgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
