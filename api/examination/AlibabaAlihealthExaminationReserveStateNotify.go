package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveStateNotify 体检机构对接_体检状态主动通知
// alibaba.alihealth.examination.reserve.state.notify
//
// 到了体检当天后，服务商主动通知体检预约状态
func AlibabaAlihealthExaminationReserveStateNotify(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveStateNotifyAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveStateNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
