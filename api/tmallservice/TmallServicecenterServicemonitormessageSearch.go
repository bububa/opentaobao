package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicemonitormessagesearch 根据时间段查询服务商的服务预警消息列表(15分钟内)
// tmall.servicecenter.servicemonitormessage.search
//
// 根据时间段查询服务商的服务预警消息列表(15分钟内)
func Tmallservicecenterservicemonitormessagesearch(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicemonitormessagesearchAPIRequest, session string) (*tmallservice.TmallservicecenterservicemonitormessagesearchAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicemonitormessagesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
