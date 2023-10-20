package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthBabyRemindBatchSend 批量发送母婴提醒
// alibaba.alihealth.baby.remind.batch.send
//
// 批量发送母婴提醒
func AlibabaAlihealthBabyRemindBatchSend(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthBabyRemindBatchSendAPIRequest, resp *alihealthcrm.AlibabaAlihealthBabyRemindBatchSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
