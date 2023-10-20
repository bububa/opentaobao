package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardServiceprogressUpdate 回传工单服务进度
// tmall.servicecenter.workcard.serviceprogress.update
//
// 回传工单服务进度
func TmallServicecenterWorkcardServiceprogressUpdate(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardServiceprogressUpdateAPIRequest, session string) (*tmallsc.TmallServicecenterWorkcardServiceprogressUpdateAPIResponse, error) {
	var resp tmallsc.TmallServicecenterWorkcardServiceprogressUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
