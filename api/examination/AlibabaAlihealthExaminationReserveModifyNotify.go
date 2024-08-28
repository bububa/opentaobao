package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveModifyNotify 通知改期结果
// alibaba.alihealth.examination.reserve.modify.notify
//
// 体检状态为改期中，服务上通知健康是否改期成功
func AlibabaAlihealthExaminationReserveModifyNotify(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
