package alihealthcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthBabyRemindBatchSend 批量发送母婴提醒
// alibaba.alihealth.baby.remind.batch.send
//
// 批量发送母婴提醒
func AlibabaAlihealthBabyRemindBatchSend(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthBabyRemindBatchSendAPIRequest, resp *alihealthcrm.AlibabaAlihealthBabyRemindBatchSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
