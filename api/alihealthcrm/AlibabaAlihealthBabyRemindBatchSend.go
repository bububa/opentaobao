package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthbabyremindbatchsend 批量发送母婴提醒
// alibaba.alihealth.baby.remind.batch.send
//
// 批量发送母婴提醒
func Alibabaalihealthbabyremindbatchsend(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthbabyremindbatchsendAPIRequest, session string) (*alihealthcrm.AlibabaalihealthbabyremindbatchsendAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthbabyremindbatchsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
