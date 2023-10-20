package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicemonitormessageupdate 服务商更新预警消息状态
// tmall.servicecenter.servicemonitormessage.update
//
// 服务商收到预警后，需要进行回复已读状态，并可填写备注
func Tmallservicecenterservicemonitormessageupdate(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicemonitormessageupdateAPIRequest, session string) (*tmallservice.TmallservicecenterservicemonitormessageupdateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicemonitormessageupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
