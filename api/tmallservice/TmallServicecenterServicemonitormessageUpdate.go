package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicemonitormessageUpdate 服务商更新预警消息状态
// tmall.servicecenter.servicemonitormessage.update
//
// 服务商收到预警后，需要进行回复已读状态，并可填写备注
func TmallServicecenterServicemonitormessageUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicemonitormessageUpdateAPIRequest, resp *tmallservice.TmallServicecenterServicemonitormessageUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
