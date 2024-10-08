package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveCancel 体检机构对接_预约取消
// alibaba.alihealth.examination.reserve.cancel
//
// 体检机构对接_体检取消
func AlibabaAlihealthExaminationReserveCancel(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveCancelAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
