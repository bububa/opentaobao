package fundplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardordersStatusMakeFinish 制卡商通知制卡完成
// alibaba.fundplatform.cardorders.status.make.finish
//
// 当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。
func AlibabaFundplatformCardordersStatusMakeFinish(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardordersStatusMakeFinishAPIRequest, resp *fundplatform.AlibabaFundplatformCardordersStatusMakeFinishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
