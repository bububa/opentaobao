package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterworkcardserviceprogressupdate 回传工单服务进度
// tmall.servicecenter.workcard.serviceprogress.update
//
// 回传工单服务进度
func Tmallservicecenterworkcardserviceprogressupdate(clt *core.SDKClient, req *tmallsc.TmallservicecenterworkcardserviceprogressupdateAPIRequest, session string) (*tmallsc.TmallservicecenterworkcardserviceprogressupdateAPIResponse, error) {
	var resp tmallsc.TmallservicecenterworkcardserviceprogressupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
