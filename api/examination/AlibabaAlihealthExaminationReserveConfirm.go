package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveConfirm 体检机构对接_体检套餐预定确认
// alibaba.alihealth.examination.reserve.confirm
//
// 向体检机构确认用户购买的体检套餐信息
func AlibabaAlihealthExaminationReserveConfirm(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveConfirmAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
