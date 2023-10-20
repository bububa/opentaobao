package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabamsfserviceauditsupdate 操作改约审批单
// alibaba.msfservice.audits.update
//
// 操作改约审批单
func Alibabamsfserviceauditsupdate(clt *core.SDKClient, req *tmallsc.AlibabamsfserviceauditsupdateAPIRequest, session string) (*tmallsc.AlibabamsfserviceauditsupdateAPIResponse, error) {
	var resp tmallsc.AlibabamsfserviceauditsupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
