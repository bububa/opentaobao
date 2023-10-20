package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardreservefail 预约失败
// tmall.servicecenter.workcard.reservefail
//
// 服务商调用该接口回传工单预约失败
func Tmallservicecenterworkcardreservefail(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardreservefailAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardreservefailAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardreservefailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
