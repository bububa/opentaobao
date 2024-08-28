package sungari

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// TaobaoSungariInspectionSubmit 抽检指令录入
// taobao.sungari.inspection.submit
//
// 抽检指令录入
func TaobaoSungariInspectionSubmit(ctx context.Context, clt *core.SDKClient, req *sungari.TaobaoSungariInspectionSubmitAPIRequest, resp *sungari.TaobaoSungariInspectionSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
