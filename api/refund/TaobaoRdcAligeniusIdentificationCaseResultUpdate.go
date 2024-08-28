package refund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRdcAligeniusIdentificationCaseResultUpdate 鉴定工单结果同步
// taobao.rdc.aligenius.identification.case.result.update
//
// 同步鉴定工单结果信息
func TaobaoRdcAligeniusIdentificationCaseResultUpdate(ctx context.Context, clt *core.SDKClient, req *refund.TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest, resp *refund.TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
