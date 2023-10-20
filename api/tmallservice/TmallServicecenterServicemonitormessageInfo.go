package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicemonitormessageinfo 服务预警单查询接口
// tmall.servicecenter.servicemonitormessage.info
//
// 服务预警单查询接口,用于查询预警单最新状态
func Tmallservicecenterservicemonitormessageinfo(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicemonitormessageinfoAPIRequest, session string) (*tmallservice.TmallservicecenterservicemonitormessageinfoAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicemonitormessageinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
